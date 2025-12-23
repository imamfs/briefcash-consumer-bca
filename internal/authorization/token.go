package authorization

import (
	"briefcash-consumer-bca/internal/entity"
	"briefcash-consumer-bca/internal/helper"
	"briefcash-consumer-bca/internal/message"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func readPemFile(fileloc string) ([]byte, error) {
	key, err := os.ReadFile(fileloc)
	if err != nil {
		return nil, fmt.Errorf("failed to read PEM file")
	}
	return key, nil
}

func loadPrivateKey(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block with private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		key2, err2 := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err2 != nil {
			return nil, err2
		}
		return key2, nil
	}

	pk, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not RSA private key")
	}
	return pk, nil
}

func signWithRsa(privateKey *rsa.PrivateKey, data string) (string, error) {
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func GetToken(partner *entity.BankPartner, log *logrus.Entry) (message.TokenResponse, error) {
	var content message.TokenResponse
	endpoint := fmt.Sprintf("%s%s", partner.BasePath, partner.TokenPath)
	timestamp := helper.FormatTime(time.Now())
	stringToSign := fmt.Sprintf("%s|%s", partner.ClientKey, timestamp)

	log.Info("Reading PEM file")
	pemFile, err := readPemFile("./resource/private_key.pem")
	if err != nil {
		log.WithError(err).Error("Failed to read PEM file")
		return message.TokenResponse{}, err
	}

	log.Info("Extract private key from PEM file")
	privateKey, err := loadPrivateKey(pemFile)
	if err != nil {
		log.WithError(err).Error("Failed to extract private key from PEM file")
		return message.TokenResponse{}, err
	}

	log.Info("Generate signature")
	signature, err := signWithRsa(privateKey, stringToSign)
	if err != nil {
		log.WithError(err).Error("Failed to generate RSA signature")
		return message.TokenResponse{}, err
	}

	payload := map[string]string{
		"grant_type": "client_credentials",
	}

	log.Info("Parsing body to byte json")
	body, err := json.Marshal(payload)
	if err != nil {
		log.WithError(err).Error("Failed to parse body to byte json")
		return message.TokenResponse{}, err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-TIMESTAMP":  timestamp,
		"X-CLIENT-KEY": partner.ClientKey,
		"X-SIGNATURE":  signature,
	}

	log.Infof("Send token request to Bank %s", partner.BankName)
	client := helper.NewHttpHelper(10 * time.Second)
	response, status, err := client.Send("POST", endpoint, body, headers)

	if err != nil {
		log.WithError(err).Errorf("Unexpected error while sending token request to bank %s", partner.BankName)
		return message.TokenResponse{}, err
	}

	if status != http.StatusOK {
		log.Errorf("Request token failed and return http status %d", status)
		return message.TokenResponse{}, fmt.Errorf("access unauthorized: http code %d", status)
	}

	log.Info("Response received, parsing byte json to struct")
	if err := json.Unmarshal(response, &content); err != nil {
		log.WithError(err).Error("Failed to parse byte json to struct")
		return message.TokenResponse{}, err
	}

	if content.AccessToken == "" {
		log.Error("Access token is empty")
		return message.TokenResponse{}, fmt.Errorf("access token is empty")
	}

	log.Info("Access token successfully retrieved")

	return content, nil
}
