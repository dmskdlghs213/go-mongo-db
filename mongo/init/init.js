const users = [
    {
      user: "root",
      pwd: "root",
      roles: [
        {
          role: "dbOwner",
          db: "test_dev"
        }
      ]
    }
];
  
  for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
  }