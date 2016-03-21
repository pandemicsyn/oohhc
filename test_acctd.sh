#! /bin/bash
# Functional test of oohhc-acctd

# List accounts
echo -e "/nLIST Action/n"
oohhc-cli -k 123456789 list | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Create account
echo -e "/nCREATE Action/n"
oohhc-cli -k 123456789 create -N company1 | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Get account
echo -e "/nGET Action/n"
oohhc-cli -k 123456789 get be1e82f6-652f-4840-97c6-aad31db6ff70 | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Update account (Name)
echo -e "/nUPDATE (Name) Action/n"
oohhc-cli -k 123456789 update be1e82f6-652f-4840-97c6-aad31db6ff70 -N newname | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Update account (Token)
echo -e "/nUPDATE (Token) Action/n"
oohhc-cli -k 123456789 update be1e82f6-652f-4840-97c6-aad31db6ff70 -T | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Update account (Status)
echo -e "/nUPDATE (Status) Action\n"
oohhc-cli -k 123456789 update be1e82f6-652f-4840-97c6-aad31db6ff70 -S suspend | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Delete account
echo -e "/nDELETE Action\n"
oohhc-cli -k 123456789 be1e82f6-652f-4840-97c6-aad31db6ff70 | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi

# Update account (Undelete Status)
echo -e "/nUPDATE (Status Undelete) Action\n"
oohhc-cli -k 123456789 update be1e82f6-652f-4840-97c6-aad31db6ff70 -S active | python -m json.tool
if [ $? -eq 0 ]
then
  echo "========================================="
else
  echo "*****************************************"
fi
