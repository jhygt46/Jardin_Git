-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost
-- Tiempo de generación: 04-02-2024 a las 01:45:16
-- Versión del servidor: 8.0.17
-- Versión de PHP: 7.3.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `jardin`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `agenda`
--

CREATE TABLE `agenda` (
  `id_age` int(4) NOT NULL,
  `fecha` date NOT NULL,
  `ultima_actualizacion` datetime NOT NULL,
  `id_usr` int(4) NOT NULL,
  `ali1` tinyint(1) NOT NULL,
  `ali2` tinyint(1) NOT NULL,
  `ali3` tinyint(1) NOT NULL,
  `dep1` tinyint(1) NOT NULL,
  `dep2` tinyint(1) NOT NULL,
  `comentario` varchar(255) COLLATE utf8_spanish2_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `agenda`
--

INSERT INTO `agenda` (`id_age`, `fecha`, `ultima_actualizacion`, `id_usr`, `ali1`, `ali2`, `ali3`, `dep1`, `dep2`, `comentario`) VALUES
(6, '2023-06-15', '2023-06-15 15:23:22', 17, 1, 2, 3, 2, 3, ''),
(7, '2023-06-15', '2023-06-15 15:23:24', 18, 1, 3, 3, 3, 4, '');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `cursos`
--

CREATE TABLE `cursos` (
  `id_cur` int(4) NOT NULL,
  `nombre` varchar(100) COLLATE utf8_spanish2_ci NOT NULL,
  `eliminado` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `cursos`
--

INSERT INTO `cursos` (`id_cur`, `nombre`, `eliminado`) VALUES
(1, 'Medio Menor', 0),
(2, 'Medio Mayor', 0),
(3, 'Sala Azul3', 0);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_online`
--

CREATE TABLE `curso_online` (
  `id_cuo` int(4) NOT NULL,
  `nombre` varchar(50) COLLATE utf8_spanish2_ci NOT NULL,
  `url` varchar(60) COLLATE utf8_spanish2_ci NOT NULL,
  `nino` tinyint(1) NOT NULL,
  `visible` tinyint(1) NOT NULL,
  `eliminado` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `curso_online`
--

INSERT INTO `curso_online` (`id_cuo`, `nombre`, `url`, `nino`, `visible`, `eliminado`) VALUES
(1, 'Cuentos', 'cuentos', 1, 1, 0),
(2, 'Cuentos Narrados', 'cuentos_narrados', 2, 1, 0),
(3, 'Canciones', 'canciones', 3, 1, 0),
(4, 'Cuento 1', 'cuento_1', 1, 1, 0),
(5, 'Cuento 1', 'cuento_11', 2, 1, 0);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_online_items`
--

CREATE TABLE `curso_online_items` (
  `id_coi` int(4) NOT NULL,
  `nombre` varchar(70) COLLATE utf8_spanish2_ci NOT NULL,
  `url` varchar(80) COLLATE utf8_spanish2_ci NOT NULL,
  `url_externo` varchar(255) COLLATE utf8_spanish2_ci NOT NULL,
  `image` varchar(70) COLLATE utf8_spanish2_ci NOT NULL,
  `lista_imagen` text COLLATE utf8_spanish2_ci NOT NULL,
  `Dx` int(4) NOT NULL,
  `Dy` int(4) NOT NULL,
  `tipo` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `curso_online_items`
--

INSERT INTO `curso_online_items` (`id_coi`, `nombre`, `url`, `url_externo`, `image`, `lista_imagen`, `Dx`, `Dy`, `tipo`) VALUES
(1, 'Te quiero tanto, MAMÁ Hola Mundo', 'te_quiero_tanto_mama', '', 'image1.jpg', '[{\"Nom\":\"1_0.jpg\"},{\"Nom\":\"1_3.jpg\"},{\"Nom\":\"1_0.jpg\"},{\"Nom\":\"1_3.jpg\"}]', 483, 400, 1),
(2, 'La cebra Camila', 'la_cebra_camila', 'CAwOO6NgbyU', 'prev1.jpg', '', 483, 400, 3),
(3, 'wegwerg', 'wegwerg', 'video.mp4', 'prev1.jpg', '', 483, 400, 2),
(4, 'Prueba', 'prueba', 'https://www.jardinvalleencantado.cl', 'prev1.jpg', '', 0, 0, 4);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_online_rel`
--

CREATE TABLE `curso_online_rel` (
  `id_cuo` int(4) NOT NULL,
  `id_coi` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `curso_online_rel`
--

INSERT INTO `curso_online_rel` (`id_cuo`, `id_coi`) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_usuarios`
--

CREATE TABLE `curso_usuarios` (
  `id_cur` int(4) NOT NULL,
  `id_usr` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `curso_usuarios`
--

INSERT INTO `curso_usuarios` (`id_cur`, `id_usr`) VALUES
(3, 14),
(3, 16),
(2, 17),
(2, 18);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `educadora_curso`
--

CREATE TABLE `educadora_curso` (
  `id_usr` int(4) NOT NULL,
  `id_cur` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `educadora_curso`
--

INSERT INTO `educadora_curso` (`id_usr`, `id_cur`) VALUES
(1, 1),
(15, 1),
(1, 2),
(15, 2),
(1, 3);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `libros`
--

CREATE TABLE `libros` (
  `id_lib` int(4) NOT NULL,
  `nombre` varchar(100) COLLATE utf8_spanish2_ci NOT NULL,
  `code` varchar(32) COLLATE utf8_spanish2_ci NOT NULL,
  `eliminado` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `libros`
--

INSERT INTO `libros` (`id_lib`, `nombre`, `code`, `eliminado`) VALUES
(1, 'Hola Mundo', 'grwthe54yw4ge4gw34g43r43f4f34f35', 0),
(2, 'osito', 'grwthe54yw4ge4gw34g43r43f4f34f34', 0),
(3, 'Papelucho', 'grwthe54yw4ge4gw34g43r43f4f34f39', 0),
(4, 'libro1', 'grwthe54yw4ge4gw34g43r43f4f34f36', 0),
(5, 'libro2', 'grwthe54yw4ge4gw34g43r43f4f34f37', 0),
(6, 'libro3', 'grwthe54yw4ge4gw34g43r43f4f34f38', 0),
(7, 'libro6', 'grwthe54yw4ge4gw34g43r43f4f34f20', 0),
(8, 'libro8', 'grwthe54yw4ge4gw34g43r43f4f34f21', 0),
(9, 'libro9', 'grwthe54yw4ge4gw34g43r43f4f34f22', 0),
(10, 'libro5', 'grwthe54yw4ge4gw34g43r43f4f34f23', 0),
(11, 'libro7', 'grwthe54yw4ge4gw34g43r43f4f34f25', 0),
(12, 'libro10', 'grwthe54yw4ge4gw34g43r43f4f34f26', 0),
(13, 'libro11', 'grwthe54yw4ge4gw34g43r43f4f34f27', 0),
(14, 'libro12', 'grwthe54yw4ge4gw34g43r43f4f34f28', 0);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `parentensco`
--

CREATE TABLE `parentensco` (
  `id_apo` int(4) NOT NULL,
  `id_alu` int(4) NOT NULL,
  `tipo` tinyint(1) NOT NULL,
  `apoderado` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `parentensco`
--

INSERT INTO `parentensco` (`id_apo`, `id_alu`, `tipo`, `apoderado`) VALUES
(1, 17, 0, 1),
(1, 18, 0, 0),
(10, 14, 0, 0),
(19, 14, 0, 0);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `prestamos`
--

CREATE TABLE `prestamos` (
  `id_pre` int(4) NOT NULL,
  `fecha_prestamo` date NOT NULL,
  `fecha_devolucion` datetime NOT NULL,
  `id_lib` int(4) NOT NULL,
  `id_alu` int(4) NOT NULL,
  `id_edu` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `prestamos`
--

INSERT INTO `prestamos` (`id_pre`, `fecha_prestamo`, `fecha_devolucion`, `id_lib`, `id_alu`, `id_edu`) VALUES
(15, '2023-06-17', '0000-00-00 00:00:00', 8, 18, 1),
(17, '2023-06-17', '0000-00-00 00:00:00', 10, 17, 1),
(18, '2023-06-17', '0000-00-00 00:00:00', 11, 14, 1),
(19, '2023-06-17', '0000-00-00 00:00:00', 12, 17, 1);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `sesiones`
--

CREATE TABLE `sesiones` (
  `id_ses` int(4) NOT NULL,
  `cookie` varchar(32) COLLATE utf8_spanish2_ci NOT NULL,
  `fecha` datetime NOT NULL,
  `id_usr` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `sesiones`
--

INSERT INTO `sesiones` (`id_ses`, `cookie`, `fecha`, `id_usr`) VALUES
(18, 'msKhKOAvwHgN6vk5qiwYSRZmhICUZShm', '2023-06-16 17:09:25', 1),
(19, 'moSOyMbrzEaVuXMzD2u04Wv3FIOt1HLZ', '2023-08-24 02:14:38', 1),
(20, 'm9It06wZ8QLFai0buc9KlmgX8UVmbrbT', '2023-09-18 18:21:51', 1),
(21, 'mWLN6rJwT5lgIzDI38uPDCarFc4RvoHB', '2023-09-27 16:32:02', 1),
(23, 'mylYwi1HVDHmGyoyjhHOdeAWid85thsm', '2023-12-28 18:51:09', 1);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuarios`
--

CREATE TABLE `usuarios` (
  `id_usr` int(4) NOT NULL,
  `nombre` varchar(255) COLLATE utf8_spanish2_ci NOT NULL,
  `correo` varchar(255) COLLATE utf8_spanish2_ci NOT NULL,
  `pass` varchar(32) COLLATE utf8_spanish2_ci NOT NULL,
  `tipo` tinyint(1) NOT NULL,
  `cant_agenda` int(4) NOT NULL,
  `telefono` varchar(14) COLLATE utf8_spanish2_ci NOT NULL,
  `nmatricula` varchar(20) COLLATE utf8_spanish2_ci NOT NULL,
  `rut` varchar(20) COLLATE utf8_spanish2_ci NOT NULL,
  `apellido1` varchar(100) COLLATE utf8_spanish2_ci NOT NULL,
  `apellido2` varchar(100) COLLATE utf8_spanish2_ci NOT NULL,
  `genero` tinyint(1) NOT NULL,
  `reglamento` tinyint(1) NOT NULL,
  `fecha_nacimiento` date NOT NULL,
  `fecha_matricula` date NOT NULL,
  `fecha_ingreso` date NOT NULL,
  `direccion` varchar(255) COLLATE utf8_spanish2_ci NOT NULL,
  `fecha_retiro` date NOT NULL,
  `motivo_retiro` tinyint(1) NOT NULL,
  `observaciones` text COLLATE utf8_spanish2_ci NOT NULL,
  `telefono2` varchar(14) COLLATE utf8_spanish2_ci NOT NULL,
  `eliminado` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `usuarios`
--

INSERT INTO `usuarios` (`id_usr`, `nombre`, `correo`, `pass`, `tipo`, `cant_agenda`, `telefono`, `nmatricula`, `rut`, `apellido1`, `apellido2`, `genero`, `reglamento`, `fecha_nacimiento`, `fecha_matricula`, `fecha_ingreso`, `direccion`, `fecha_retiro`, `motivo_retiro`, `observaciones`, `telefono2`, `eliminado`) VALUES
(1, 'Eliana2', 'elibruzzo@hotmail.com', 'd19ed8f8ac7e5cd3a51a58c3511e6ea4', 0, 0, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '+56966166924', 0),
(10, 'Padre1', 'padre1@gmail.com', '', 2, 0, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(14, 'Diego', '', '', 3, 0, '', 'erherh', '15.935.774-0', 'Gomez', 'Bezmalinovic', 1, 1, '2023-06-17', '2023-06-17', '2023-06-23', 'Av. 11 de Setiembre 2155', '2023-06-17', 2, 'ergergerg', '', 0),
(15, 'Juanita', 'juanita@gmail.com', '', 1, 0, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(16, 'Diego2', '', '', 3, 0, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(17, 'Diego 3', '', '', 3, 0, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(18, 'Diego 4', '', '', 3, 0, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(19, 'Padre2', 'user10@gmail.com', '', 2, 0, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0);

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `agenda`
--
ALTER TABLE `agenda`
  ADD PRIMARY KEY (`id_age`),
  ADD KEY `id_alu` (`id_usr`);

--
-- Indices de la tabla `cursos`
--
ALTER TABLE `cursos`
  ADD PRIMARY KEY (`id_cur`);

--
-- Indices de la tabla `curso_online`
--
ALTER TABLE `curso_online`
  ADD PRIMARY KEY (`id_cuo`);

--
-- Indices de la tabla `curso_online_items`
--
ALTER TABLE `curso_online_items`
  ADD PRIMARY KEY (`id_coi`);

--
-- Indices de la tabla `curso_online_rel`
--
ALTER TABLE `curso_online_rel`
  ADD PRIMARY KEY (`id_cuo`,`id_coi`),
  ADD KEY `id_coi` (`id_coi`);

--
-- Indices de la tabla `curso_usuarios`
--
ALTER TABLE `curso_usuarios`
  ADD PRIMARY KEY (`id_cur`,`id_usr`),
  ADD KEY `id_usr` (`id_usr`);

--
-- Indices de la tabla `educadora_curso`
--
ALTER TABLE `educadora_curso`
  ADD PRIMARY KEY (`id_usr`,`id_cur`),
  ADD KEY `id_cur` (`id_cur`);

--
-- Indices de la tabla `libros`
--
ALTER TABLE `libros`
  ADD PRIMARY KEY (`id_lib`);

--
-- Indices de la tabla `parentensco`
--
ALTER TABLE `parentensco`
  ADD PRIMARY KEY (`id_apo`,`id_alu`),
  ADD KEY `id_user2` (`id_alu`);

--
-- Indices de la tabla `prestamos`
--
ALTER TABLE `prestamos`
  ADD PRIMARY KEY (`id_pre`),
  ADD KEY `id_alu` (`id_alu`),
  ADD KEY `id_edu` (`id_edu`),
  ADD KEY `id_lib` (`id_lib`);

--
-- Indices de la tabla `sesiones`
--
ALTER TABLE `sesiones`
  ADD PRIMARY KEY (`id_ses`),
  ADD KEY `id_usr` (`id_usr`);

--
-- Indices de la tabla `usuarios`
--
ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`id_usr`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `agenda`
--
ALTER TABLE `agenda`
  MODIFY `id_age` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT de la tabla `cursos`
--
ALTER TABLE `cursos`
  MODIFY `id_cur` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT de la tabla `curso_online`
--
ALTER TABLE `curso_online`
  MODIFY `id_cuo` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT de la tabla `curso_online_items`
--
ALTER TABLE `curso_online_items`
  MODIFY `id_coi` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT de la tabla `libros`
--
ALTER TABLE `libros`
  MODIFY `id_lib` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT de la tabla `prestamos`
--
ALTER TABLE `prestamos`
  MODIFY `id_pre` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT de la tabla `sesiones`
--
ALTER TABLE `sesiones`
  MODIFY `id_ses` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT de la tabla `usuarios`
--
ALTER TABLE `usuarios`
  MODIFY `id_usr` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `agenda`
--
ALTER TABLE `agenda`
  ADD CONSTRAINT `agenda_ibfk_1` FOREIGN KEY (`id_usr`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `curso_online_rel`
--
ALTER TABLE `curso_online_rel`
  ADD CONSTRAINT `curso_online_rel_ibfk_1` FOREIGN KEY (`id_cuo`) REFERENCES `curso_online` (`id_cuo`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `curso_online_rel_ibfk_2` FOREIGN KEY (`id_coi`) REFERENCES `curso_online_items` (`id_coi`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `curso_usuarios`
--
ALTER TABLE `curso_usuarios`
  ADD CONSTRAINT `curso_usuarios_ibfk_1` FOREIGN KEY (`id_cur`) REFERENCES `cursos` (`id_cur`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `curso_usuarios_ibfk_2` FOREIGN KEY (`id_usr`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `educadora_curso`
--
ALTER TABLE `educadora_curso`
  ADD CONSTRAINT `educadora_curso_ibfk_1` FOREIGN KEY (`id_cur`) REFERENCES `cursos` (`id_cur`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `educadora_curso_ibfk_2` FOREIGN KEY (`id_usr`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `parentensco`
--
ALTER TABLE `parentensco`
  ADD CONSTRAINT `parentensco_ibfk_1` FOREIGN KEY (`id_apo`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `parentensco_ibfk_2` FOREIGN KEY (`id_alu`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `prestamos`
--
ALTER TABLE `prestamos`
  ADD CONSTRAINT `prestamos_ibfk_1` FOREIGN KEY (`id_edu`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `prestamos_ibfk_2` FOREIGN KEY (`id_alu`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `prestamos_ibfk_3` FOREIGN KEY (`id_lib`) REFERENCES `libros` (`id_lib`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `sesiones`
--
ALTER TABLE `sesiones`
  ADD CONSTRAINT `sesiones_ibfk_1` FOREIGN KEY (`id_usr`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
