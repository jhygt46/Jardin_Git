-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost
-- Tiempo de generación: 15-06-2023 a las 17:33:15
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
(6, '2023-06-15', '2023-06-15 02:16:03', 17, 3, 3, 1, 3, 4, ''),
(7, '2023-06-15', '2023-06-15 02:28:26', 18, 1, 3, 3, 3, 2, '');

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
(2, 'osito', 'grwthe54yw4ge4gw34g43r43f4f34f34', 0);

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
(1, 18, 0, 0);

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
(7, '2023-06-12', '0000-00-00 00:00:00', 1, 14, 10),
(8, '2023-06-15', '0000-00-00 00:00:00', 2, 16, 1);

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
(14, 'mz5vdciGQU3BB0tvWj64yVc6NfYokb48', '2023-06-15 01:08:32', 1);

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

INSERT INTO `usuarios` (`id_usr`, `nombre`, `correo`, `pass`, `tipo`, `telefono`, `nmatricula`, `rut`, `apellido1`, `apellido2`, `genero`, `reglamento`, `fecha_nacimiento`, `fecha_matricula`, `fecha_ingreso`, `direccion`, `fecha_retiro`, `motivo_retiro`, `observaciones`, `telefono2`, `eliminado`) VALUES
(1, 'User4', 'elibruzzo@hotmail.com', '25d55ad283aa400af464c76d713c07ad', 0, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '+56966166924', 0),
(10, 'Padre1', 'padre1@gmail.com', '', 2, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(14, 'Diego', '', '', 3, '', 'erherh', '15.935.774-0', 'Gomez', 'Bezmalinovic', 1, 1, '2023-06-17', '2023-06-17', '2023-06-23', 'Av. 11 de Setiembre 2155', '2023-06-17', 2, 'ergergerg', '', 0),
(15, 'Juanita', 'juanita@gmail.com', '', 1, '+56966166923', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(16, 'Diego2', '', '', 3, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(17, 'Diego 3', '', '', 3, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0),
(18, 'Diego 4', '', '', 3, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0);

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
  MODIFY `id_cur` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT de la tabla `libros`
--
ALTER TABLE `libros`
  MODIFY `id_lib` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT de la tabla `prestamos`
--
ALTER TABLE `prestamos`
  MODIFY `id_pre` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT de la tabla `sesiones`
--
ALTER TABLE `sesiones`
  MODIFY `id_ses` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT de la tabla `usuarios`
--
ALTER TABLE `usuarios`
  MODIFY `id_usr` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `agenda`
--
ALTER TABLE `agenda`
  ADD CONSTRAINT `agenda_ibfk_1` FOREIGN KEY (`id_usr`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE;

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
