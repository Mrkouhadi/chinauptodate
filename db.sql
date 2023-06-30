-- phpMyAdmin SQL Dump
-- version 4.1.12
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: May 16, 2020 at 12:45 PM
-- Server version: 5.6.16
-- PHP Version: 5.5.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `inet_newsportal`
--

-- --------------------------------------------------------

--
-- Table structure for table `tblauthor`
--

CREATE TABLE IF NOT EXISTS `tblauthor` (
  `author_id` int(11) NOT NULL AUTO_INCREMENT,
  `author_name` varchar(50) NOT NULL,
  `author_display_name` varchar(30) NOT NULL,
  `author_email` varchar(30) NOT NULL,
  `author_account_status` int(1) NOT NULL,
  `author_profile` longblob NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(30) NOT NULL,
  PRIMARY KEY (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `tblbackup`
--

CREATE TABLE IF NOT EXISTS `tblbackup` (
  `backup_id` int(11) NOT NULL AUTO_INCREMENT,
  `backup_date` date NOT NULL,
  `backup_file` varchar(50) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`backup_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `tblcomment`
--

CREATE TABLE IF NOT EXISTS `tblcomment` (
  `comment_id` int(11) NOT NULL AUTO_INCREMENT,
  `comment_content` varchar(100) NOT NULL,
  `subscriber_id` int(11) NOT NULL,
  `news_id` int(11) NOT NULL,
  `comment_date` date NOT NULL,
  `comment_status` int(1) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`comment_id`),
  KEY `subscriber_id` (`subscriber_id`,`user_id`),
  KEY `news_id` (`news_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `tblnews`
--

CREATE TABLE IF NOT EXISTS `tblnews` (
  `news_id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` int(11) NOT NULL,
  `date_posted` date NOT NULL,
  `news_title` varchar(100) NOT NULL,
  `news_content` text NOT NULL,
  `date_updated` date NOT NULL,
  `news_status` int(1) NOT NULL,
  `comment_status` int(1) NOT NULL,
  `author_id` int(11) NOT NULL,
  PRIMARY KEY (`news_id`),
  KEY `category_id` (`category_id`,`author_id`),
  KEY `author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `tblnewscategory`
--

CREATE TABLE IF NOT EXISTS `tblnewscategory` (
  `category_id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(30) NOT NULL,
  `category_description` varchar(50) NOT NULL,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `tblsubscriber`
--

CREATE TABLE IF NOT EXISTS `tblsubscriber` (
  `subscriber_id` int(11) NOT NULL AUTO_INCREMENT,
  `subscriber_display_name` varchar(30) NOT NULL,
  `subscriber_name` varchar(50) NOT NULL,
  `subscriber_email` varchar(30) NOT NULL,
  `subscriber_profile` longblob NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(30) NOT NULL,
  `account_status` int(1) NOT NULL,
  `date_joined` date NOT NULL,
  `date_approved` date NOT NULL,
  PRIMARY KEY (`subscriber_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `tbluser`
--

CREATE TABLE IF NOT EXISTS `tbluser` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_display_name` varchar(30) NOT NULL,
  `user_complete_name` varchar(30) NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(30) NOT NULL,
  `user_profile` longblob NOT NULL,
  `user_type` int(1) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tblbackup`
--
ALTER TABLE `tblbackup`
  ADD CONSTRAINT `tblbackup_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `tbluser` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `tblcomment`
--
ALTER TABLE `tblcomment`
  ADD CONSTRAINT `tblcomment_ibfk_1` FOREIGN KEY (`subscriber_id`) REFERENCES `tblsubscriber` (`subscriber_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tblcomment_ibfk_2` FOREIGN KEY (`news_id`) REFERENCES `tblnews` (`news_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tblcomment_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `tbluser` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `tblnews`
--
ALTER TABLE `tblnews`
  ADD CONSTRAINT `tblnews_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `tblnewscategory` (`category_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tblnews_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `tblauthor` (`author_id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
