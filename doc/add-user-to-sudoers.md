# Add user to sudoers

Login as sudo
```sh
su -
```

add user to sudoers
```sh
usermod -aG sudo $USERNAME
```

reboot
```sh
reboot now
```