# Altitude Test Task

Written in:
- Go with Echo framework
- JavaScript with React library

Disclaimer:
- The code for sending the mail has been commented out because SendGrid wasn't working, so the other
smtp server was sending mails only to my registered mail. 2FA and Verification links are printed in the console.
- Nginx has an error on linux where `host.docker.internal` won't resolve, I didn't have time to fix it. If it persists,
change the `environment` file on frontend to route API calls to `:8080` instead of `:6789`. Nginx will still serve images.


# Features

### General
- [x] Login
- [x] Register

###  User
- [x] Update profile
- [x] Update image
- [x] Update password
- [x] Notification if account deleted

### Admin
- [x] Added through code
- [x] See user list
- [x] Filter by email and birthday
- [x] Delete user

### Bonus
- [x] Login with Google
- [x] Verify via Email after login
- [x] Admin filter by verification status
- [x] 2FA via Email
- [x] Enable/Disable 2FA
