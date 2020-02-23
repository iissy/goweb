CREATE SCHEMA `hrefsdb` ;

use hrefsdb;

CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` varchar(30) NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(36) NOT NULL,
  `regdate` datetime NOT NULL,
  `lastlogindate` datetime DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `account` VALUES (1,'jimmy','jimmy','e10adc3949ba59abbe56e057f20f883e','2019-01-07 09:01:19','2020-02-18 09:47:16',0);

CREATE TABLE `article` (
  `id` varchar(100) NOT NULL,
  `title` varchar(100) NOT NULL,
  `icon` varchar(200) NOT NULL,
  `visited` int(11) NOT NULL DEFAULT '0',
  `brief` varchar(500) NOT NULL,
  `body` text NOT NULL,
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `cuslink` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL,
  `url` varchar(100) NOT NULL,
  `status` int(11) NOT NULL,
  `catid` varchar(50) DEFAULT NULL,
  `visited` int(11) DEFAULT '0',
  `linktype` varchar(50) DEFAULT NULL,
  `adddate` datetime NOT NULL,
  `updatedate` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8;

CREATE TABLE `link` (
  `id` varchar(50) NOT NULL,
  `icon` varchar(200) DEFAULT NULL,
  `linktype` varchar(50) NOT NULL,
  `catid` varchar(50) DEFAULT NULL,
  `url` varchar(200) NOT NULL,
  `title` varchar(100) NOT NULL,
  `brief` varchar(1000) NOT NULL,
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `visited` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `linkcat` (
  `id` varchar(50) NOT NULL,
  `catname` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `linkcat` VALUES ('ae45b8c6-f0ed-11e9-81b7-0242ac120004','人工智能'),('ae45baa0-f0ed-11e9-81b7-0242ac120004','JQuery'),('ae45bae8-f0ed-11e9-81b7-0242ac120004','NodeJS'),('ae45bcb1-f0ed-11e9-81b7-0242ac120004','其他'),('ae45bce3-f0ed-11e9-81b7-0242ac120004','教程'),('ae45bd16-f0ed-11e9-81b7-0242ac120004','PHP'),('ae45bd20-f0ed-11e9-81b7-0242ac120004','前端框架'),('ae45bd3e-f0ed-11e9-81b7-0242ac120004','运维工具'),('ae45bd82-f0ed-11e9-81b7-0242ac120004','Vue'),('ae45bd8d-f0ed-11e9-81b7-0242ac120004','数据库'),('ae45bd97-f0ed-11e9-81b7-0242ac120004','协同工具'),('ae45bdac-f0ed-11e9-81b7-0242ac120004','大数据'),('ae45bdd8-f0ed-11e9-81b7-0242ac120004','区块链'),('ae45bdf7-f0ed-11e9-81b7-0242ac120004','Charts'),('ae45be34-f0ed-11e9-81b7-0242ac120004','软件'),('ae45be3e-f0ed-11e9-81b7-0242ac120004','富文本编辑器'),('ae45be5d-f0ed-11e9-81b7-0242ac120004','Python'),('ae45be67-f0ed-11e9-81b7-0242ac120004','Java'),('ae45be7c-f0ed-11e9-81b7-0242ac120004','DotNet'),('ae45be91-f0ed-11e9-81b7-0242ac120004','打包构建'),('ae45bea6-f0ed-11e9-81b7-0242ac120004','CSS'),('ae45beba-f0ed-11e9-81b7-0242ac120004','Golang'),('ae45becf-f0ed-11e9-81b7-0242ac120004','架构师'),('ae45bee3-f0ed-11e9-81b7-0242ac120004','实用工具'),('ae45beed-f0ed-11e9-81b7-0242ac120004','公共'),('ae45bef7-f0ed-11e9-81b7-0242ac120004','资讯');