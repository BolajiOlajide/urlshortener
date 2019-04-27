CREATE DATABASE IF NOT EXISTS demodb;

GRANT ALL PRIVILEGES ON demodb.* to bolaji@localhost;

USE demodb;
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER NOT NULL AUTO_INCREMENT,
    path VARCHAR(15) NOT NULL,
    url VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO urls(path, url) VALUES
    ('/bhgisd', 'https://facebook.com'),
    ('/astaxie', 'https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.2.html'),
    ('/svg', 'https://www.sarasoueidan.com/blog/svg-coordinate-systems/'),
    ('/turing', 'https://github.com/BolajiOlajide/turing-tshirt/tree/master/backend');
