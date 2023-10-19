package object_encrypter

import (
	"testing"
)

var instance Encrypter

func init() {
	instance = NewEncrypter("X152yX4gZcpO1xyX95W5na75lIbmVcclxuIl0sIm1hc4D0gZW5j")
}

type TestingStruct struct {
	Data struct {
		Cards []struct {
			Id         string `json:"id"`
			CardNumber string `json:"card_number"`
			IsNBU      bool   `json:"is_nbu"`
			Expire     string `json:"expire"`
			IsMain     string `json:"is_main"`
		} `json:"cards'`
	} `json:"data"`
	State     string `json:"state"`
	ErrorText string `json:"error_text"`
}

func TestDummyDecrypt(t *testing.T) {
	var c TestingStruct
	enc := BASE64
	err := instance.Decrypt("MqfXQg9frMKNJ2qoZ3vccDqqPDKqwTji1Zz+6Js5IlCdRjLVFkO9NkGKMiryTsf6zZLRqSkT7pYsKBERRfV6q5Z+97yMpdgYjgyycF39XoT81pKIzsVt09htDkgpdsWQYRkubKB7bqS8awfJv6Y7APfx5ZHCcGHRdE2SNbg0T+pmx8tZWhZfPF8ygk1s9g1Q2xYx7D4adUJIPavcUfQE6RTOcCLoXKpHlZxBoq0LfpASk9xEC+u9OHHj8ndjJBCwDOBE89Wlv7Ltp3BPVKvfBQLJlYCvZRy72Pt6Fv0NuITsY9nbF0pLXArVgMqVX3vtpx+u/40uJkiSPijSma1q550dqYE4qKy8NBXOyQBaj0b5F3Vq6diKK0ImtTF1KyyK5ZxWFd7TGLOaEKBSEOAStC+65ruf05HDGamg6GmhQTGCajC32rQ78SehSQRlRKpQ+AXbZ4xbwlpp2vAxS6AJcRoASLSKL4SXzoRiJhfIzgHwN70FFFPFsYeIgLQmJRJRxBS9mYzJc0kOZE8iWGXt5bQHyXS9cKv/ctf/h8xj2teAAKEsmwh5RXFC6ZBhHyH6dJtT28gfCaLl7SiKC2qZzXNSiq3tV0W66ijGYtW4vxpq1XbKQi6f/2UgH3H5kJb70P82K6VHROvW2VHPdcZScJzcZg5QmM6J5f4kQisNhkT2cnpOmat4aKCIRulUn8P/A5TCyfy3G4Ii1oS+RwwJ/sz07MNRhdN8HcyJ5ceT4P5x/MZF5TXN7oMMvKhLMSwfmfy7BYERUmtziUHETDahLYLHsGVxjtDPD15iQcLlsRodL5tWAzZPNZlkw5fpPoDSAGxYA+7g5vfBkutY+jkCSI6ukV4BaZulRtyZ/RzIvHTje2adF5aEJ+VzOFm3+zKBJ+oBoTdekP9kpPBOp4cI+fuQr8qssJvi1utH+RTxdvRScFYzMI99peZsTXJ3Fnjs4RrHiE5TrbUYgrsJ1CeXSAnDT78N6tKV30BOCK7mymavg/PLEWvODjAEUGQLqgW2B7xOaX+BzQgIkBYzr32zW6HoDS/XDPlC+6ziatwBDChSSKnXITrxqrMwJZeS0ihn5twPmN1o5rUmdBZGXiZEap4kqTFuZ6WdZPJzzPO9EFI=", &c, &enc)
	if err != nil {
		t.Error(
			"Expected", "no error",
			"Got", err.Error(),
		)
		t.FailNow()
	}
	if c.State != "success" {
		t.Error(
			"Expected", "the state of struct must be success",
			"Got", c.State,
		)
	}
}
