<!--
page_title: J-USB, My USB Password Setup
page_description: I've started moving away from online password storage, and started using KeePassXC on a USB stick; here's how my setup works...
page_status: published
page_date: 2022/02/07
page_image: https://julianorchard.co.uk/res/default.png
-->

# My Offline Password Setup

I've recently, finally, started to move away from the standard online Google
password storage system. This is for a few reasons, but it's just something I've
never felt 100% comfortable with. 

The best option I found that I wanted to briefly write about was the open source
[KeePassXC](https://github.com/keepassxreboot/keepassxc). The
[portable](https://github.com/keepassxreboot/keepassxc/releases/tag/2.6.6)
version of the project is what I found to be the best, as I wanted to be able to
manage passwords even on machines I'd not visited, and figured the best way to
do this, offline, would be to just carry them on a USB Stick. 

## Backing Up The Database

KeePassXC generates a database of passwords, `Passwords.kdbx` as standard. This
is a script I wrote to back this up, so that when I plug in the USB into my
home machine, I can just double click this script to run it; 

```
' File:        backup.vbs
' Author:      Julian Orachrd (hello@julianorchard.co.uk)
' Description: This script mainly exists to ensure I
'              remember to consider if a machine is
'              trustworthy before backing up the database file
' Date:        10/01/2022

' File System Object
	Set fso = CreateObject("Scripting.FileSystemObject")
' Backup Folder
	bkp = CreateObject("WScript.Shell").ExpandEnvironmentStrings("%USERPROFILE%") & "\JUSBackup\"
	bkf = bkp & "Passwords.kdbx"

' If there's already a backup database file, just proceed
	If fso.FileExists(bkf) Then
		proceed=6
	Else
		proceed=MsgBox("Are you SURE you want to back up your KeePassXC Database to this machine?" & vbNewLine & vbNewLine & "Only back up the database to a machine you trust." & vbNewLine & vbNewLine & "Saves to: " & bkp, 308, "Trust No One, Not Even Your Printer")

	End If

' Copy database if either user is happy to do so
' or if the database file already exists...
	If proceed = 6 Then
		If Not fso.FolderExists(bkp) Then
			fso.CreateFolder bkp
		End If
		fso.CopyFile "KeePassXC\Passwords.kdbx", bkf
	End If
```

It's mainly to remind me not to back up the password database to machines I
don't trust...

That's it, just wanted to share my setup as I've had fun doing it.
