-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: May 02, 2020 at 01:29 AM
-- Server version: 5.7.23
-- PHP Version: 7.2.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `web_gudang`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `ambil` ()  BEGIN
    SELECT nama_barang FROM outcoming;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `ambilNama` ()  BEGIN
    SELECT namabarang FROM outcoming;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `selectNamaBarang` ()  BEGIN
    SELECT nama_barang FROM outcoming;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `incoming`
--

CREATE TABLE `incoming` (
  `id` int(11) NOT NULL,
  `transaksi_id` varchar(50) NOT NULL,
  `tanggal` varchar(20) NOT NULL,
  `lokasi` varchar(100) NOT NULL,
  `kode_barang` varchar(100) NOT NULL,
  `nama_barang` varchar(100) NOT NULL,
  `satuan` varchar(50) NOT NULL,
  `jumlah` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `incoming`
--

INSERT INTO `incoming` (`id`, `transaksi_id`, `tanggal`, `lokasi`, `kode_barang`, `nama_barang`, `satuan`, `jumlah`) VALUES
(126, 'TRX-130-724', '2020/05/02 08:23:17', 'Makasar', 'MKS', 'Kain Garmen', 'KODI', 432),
(127, 'TRX-294-759', '2020/05/02 08:24:09', 'Surabaya', 'SBY', 'Boeing 737', 'Unit', 2);

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `item_name` varchar(100) NOT NULL,
  `item_code` varchar(100) NOT NULL,
  `unit_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `outcoming`
--

CREATE TABLE `outcoming` (
  `id` int(10) NOT NULL,
  `transaksi_id` varchar(50) NOT NULL,
  `tanggal_masuk` varchar(20) NOT NULL,
  `tanggal_keluar` varchar(20) NOT NULL,
  `lokasi` varchar(100) NOT NULL,
  `kode_barang` varchar(100) NOT NULL,
  `nama_barang` varchar(100) NOT NULL,
  `satuan` varchar(50) NOT NULL,
  `jumlah` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `outcoming`
--

INSERT INTO `outcoming` (`id`, `transaksi_id`, `tanggal_masuk`, `tanggal_keluar`, `lokasi`, `kode_barang`, `nama_barang`, `satuan`, `jumlah`) VALUES
(17, 'PTRO-TBG-201980963215', '31/12/2019', '31/12/2019', 'PTRO-TABANG', 'TBG-291029', 'Jigsaw', 'Dus', '1');

--
-- Triggers `outcoming`
--
DELIMITER $$
CREATE TRIGGER `TG_BARANG_KELUAR` AFTER INSERT ON `outcoming` FOR EACH ROW BEGIN
 UPDATE tb_barang_masuk SET jumlah=jumlah-NEW.jumlah
 WHERE kode_barang=NEW.kode_barang;
 DELETE FROM tb_barang_masuk WHERE jumlah = 0;

END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `tb_upload_gambar_user`
--

CREATE TABLE `tb_upload_gambar_user` (
  `id` int(11) NOT NULL,
  `username_user` varchar(100) NOT NULL,
  `nama_file` varchar(220) NOT NULL,
  `ukuran_file` varchar(8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tb_upload_gambar_user`
--

INSERT INTO `tb_upload_gambar_user` (`id`, `username_user`, `nama_file`, `ukuran_file`) VALUES
(1, 'test', 'nopic5.png', '6.33'),
(2, 'test', 'nopic4.png', '6.33'),
(3, 'coba', 'ptro.jpg', '16.69'),
(4, 'admin', 'Untitled4.png', '232.91'),
(5, 'danyismail', 'nopic2.png', '6.33');

-- --------------------------------------------------------

--
-- Stand-in structure for view `total_barang_masuk`
-- (See below for the actual view)
--
CREATE TABLE `total_barang_masuk` (
`count(1)` bigint(21)
);

-- --------------------------------------------------------

--
-- Table structure for table `units`
--

CREATE TABLE `units` (
  `id` int(11) NOT NULL,
  `code` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `units`
--

INSERT INTO `units` (`id`, `code`, `name`) VALUES
(1, 'Dus', 'Dus'),
(2, 'Pcs', 'Pieces'),
(5, 'Pack', 'Pack');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(12) NOT NULL,
  `username` varchar(200) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(200) NOT NULL,
  `role` tinyint(4) NOT NULL DEFAULT '0',
  `last_login` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `role`, `last_login`) VALUES
(20, 'admin', 'admin@gmail.com', '$2y$10$3HNkMOtwX8X88Xb3DIveYuhXScTnJ9m16/rPDF1/VTa/VTisxVZ4i', 1, '27-04-2020 19:05'),
(21, 'haris', 'harisds94@gmail.com', '$2y$10$wj7T1eyikr9NTFrGVrRPQ./mt9UZgOoUshlbxOKEBKOg91j4X9D4W', 1, '31-12-2019 15:50'),
(22, 'arman', 'armanbudimahendra@gmail.com', '$2y$10$ZW.aM8FUZjCQXJWZJV5hhuUdm6/wmzqDeygOb1kBKu8NaO5zWRWru', 1, ''),
(23, 'user', 'user@gmail.com', '$2y$10$Cpab9tJemKXPTfwZBTgkFeRo8DmnV26NyXs5fiz25.Je2Aj8Kmgmq', 0, '');

-- --------------------------------------------------------

--
-- Structure for view `total_barang_masuk`
--
DROP TABLE IF EXISTS `total_barang_masuk`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `total_barang_masuk`  AS  select count(1) AS `count(1)` from `incoming` ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `incoming`
--
ALTER TABLE `incoming`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `outcoming`
--
ALTER TABLE `outcoming`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tb_upload_gambar_user`
--
ALTER TABLE `tb_upload_gambar_user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `units`
--
ALTER TABLE `units`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `incoming`
--
ALTER TABLE `incoming`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=128;

--
-- AUTO_INCREMENT for table `items`
--
ALTER TABLE `items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `outcoming`
--
ALTER TABLE `outcoming`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `tb_upload_gambar_user`
--
ALTER TABLE `tb_upload_gambar_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `units`
--
ALTER TABLE `units`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(12) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;
