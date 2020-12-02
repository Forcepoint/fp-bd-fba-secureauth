package structs

var EventTypes = map[string]string {
	"21090": "Invalid username/password.",
	"24100": "PIN Verification Successful.",
	"24110": "KBA Verification Successful",
	"24120": "OTP Verification Successful",
	"24200": "Incorrect PIN number entered.",
	"24210": "Incorrect KBA attempt.",
	"24220": "Incorrect OTP attempt.",
	"51160": "Password could not be validated.",
	"51170": "Password validated. Successful login.",
}

var EventSuccess = map[string]bool {
	"21090": false,
	"24100": false,
	"24110": false,
	"24120": false,
	"24200": false,
	"24210": false,
	"24220": false,
	"51160": false,
	"51170": true,
}