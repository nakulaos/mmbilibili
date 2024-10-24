package tools

import "strings"

func MaskPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}

	return phone[:3] + "****" + phone[len(phone)-4:]
}

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}
	username := parts[0]
	domain := parts[1]

	if len(username) <= 2 {
		return email
	}

	maskedUsername := username[:1] + "****" + username[len(username)-1:]

	return maskedUsername + "@" + domain
}
