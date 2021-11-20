# Account Management Console Application

## Basic syntax
```bash
acc [Subcommand] [Options] [Params]
```

## Prerequisite

go 1.17

## How to build
```bash
go install
```

## ADD
**add** adds account information into database.
1. Options
    * -t    : Title
    * -u    : User ID
    * -p    : Password
    * -l    : Location (URL)
    * -e    : E-mail
    * -m    : Memo
    * -a    : alias for the account
    * -f    : Insert data from json file
    > You can omit options.
    > **acc** will prompt required user inputs.
    > You don't need to give every input.

2. Params
    * None

3. Examples
```bash
acc add
```
Above command will ask you every informations to register account record prompting on console.

```bash
acc add -t title -u userid -p password
```
Above command will ask you remaining informations URL, e-mail, alias, memo prompting on console. In case of alias, you can give multiple alias with whitespace.


## DEL
**del** deletes account information from database.
1. Options
    * -t    : Title
    * -u    : User ID
    * -p    : Password
    * -l    : Location (URL)
    * -e    : E-mail
    * -m    : Memo
    * -a    : Alias
    > If you pass any option, **del** will empty passed fields of records rather than delete records themselves.
    > In case of **-a**, **del** will remove all elements of existing alias array.

2. Params
    * Index
    > You can give multiple indexs to delete with whitespace.

3. Examples
```bash
acc del 2 5 14
```
Above command will delete account informations which have 2 or 5 or 14 index number.

```bash
acc del -e -a 2 5
```
Above command will empty email and alias field of account informations which have 2 or 5 index number.


## LS
**ls** lists account information
1. Options
    * -o    : Output format
    * -i    : Lists accounts matched to the index
    * -t    : Lists accounts matched to the title
    * -u    : Lists accounts matched to the user ID
    * -a    : Lists accounts containing the alias
    > If you don't give any options for **ls** command, **ls** searches all records and prints **short** format by default.
    > Each searching option is combined by **$and** operator.

2. Output formats
    * short             : Only prints index, title, user ID, password on oneline
    * wide              : Prints every columns on multiple lines
    * format=[text]     : Prints customized format
        * %i    : Index
        * %t    : Title
        * %u    : User ID
        * %p    : Password
        * %U    : URL
        * %e    : E-mail
        * %a    : Alias
        * %m    : Memo

3. Params
    * None

4. Examples
```bash
<Command>
    acc ls -o wide

<Result>
    Index:  4
    Title:  Github
    Uid:    taeho911
    Pwd:    dummyPwd123
    Url:    github.com
    Email:  kim911@gmail.com
    Alias:  [dev git repo]
    Memo:   
    ...
```

```bash
<Command>
    acc ls -o short -t gmail

<Result>
    INDEX   TITLE       UID         PWD
    7       gmail-jp    kim911      dummyPwd123
    13      gmail       hwangdal    aawe0099
```

```bash
<Command>
    acc ls -o "format:This is customized format %i %t" -t gmail

<Result>
    This is customized format 7 gmail-jp
    This is customized format 13 gmail
```

## MOD
**mod** modifies fields of existing account information
1. Options
    * -t     : Title
    * -u     : User ID
    * -p     : Password
    * -l     : Location (URL)
    * -e     : E-mail
    * -m     : Memo
    * -a     : Alias
    * --push : Push alias into existing alias array
    * --pull : Pull alias from existing alias array
    > Regarding alias modification, **acc mod** pushes given alias into existing alias array by default.

2. Params
    * Index
    > You can give multiple indexs to delete with whitespace.
