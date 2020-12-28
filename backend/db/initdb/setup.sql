CREATE USER 'studyapp'@'%' IDENTIFIED BY 'studyapp_pw';
GRANT ALL PRIVILEGES ON studyapp_development.* TO 'studyapp'@'%';
GRANT ALL PRIVILEGES ON studyapp_test.* TO 'studyapp'@'%';
FLUSH PRIVILEGES;
