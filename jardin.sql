-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost
-- Tiempo de generación: 05-02-2024 a las 05:02:16
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
  `fecha` date NOT NULL DEFAULT '0000-00-00',
  `ultima_actualizacion` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `id_usr` int(4) NOT NULL DEFAULT '0',
  `ali1` tinyint(1) NOT NULL DEFAULT '0',
  `ali2` tinyint(1) NOT NULL DEFAULT '0',
  `ali3` tinyint(1) NOT NULL DEFAULT '0',
  `dep1` tinyint(1) NOT NULL DEFAULT '0',
  `dep2` tinyint(1) NOT NULL DEFAULT '0',
  `comentario` varchar(255) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `cursos`
--

CREATE TABLE `cursos` (
  `id_cur` int(4) NOT NULL,
  `nombre` varchar(100) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `eliminado` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_online`
--

CREATE TABLE `curso_online` (
  `id_cuo` int(4) NOT NULL,
  `nombre` varchar(50) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `url` varchar(60) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `nino` tinyint(1) NOT NULL DEFAULT '0',
  `visible` tinyint(1) NOT NULL DEFAULT '0',
  `eliminado` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_online_items`
--

CREATE TABLE `curso_online_items` (
  `id_coi` int(4) NOT NULL,
  `nombre` varchar(70) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `url` varchar(80) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `url_externo` varchar(255) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `image` varchar(70) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `lista_imagen` text CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL,
  `Dx` int(4) NOT NULL DEFAULT '0',
  `Dy` int(4) NOT NULL DEFAULT '0',
  `tipo` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_online_rel`
--

CREATE TABLE `curso_online_rel` (
  `id_cuo` int(4) NOT NULL,
  `id_coi` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `curso_usuarios`
--

CREATE TABLE `curso_usuarios` (
  `id_cur` int(4) NOT NULL,
  `id_usr` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `educadora_curso`
--

CREATE TABLE `educadora_curso` (
  `id_usr` int(4) NOT NULL,
  `id_cur` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `libros`
--

CREATE TABLE `libros` (
  `id_lib` int(4) NOT NULL,
  `nombre` varchar(100) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `code` varchar(32) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `eliminado` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `parentensco`
--

CREATE TABLE `parentensco` (
  `id_apo` int(4) NOT NULL,
  `id_alu` int(4) NOT NULL,
  `tipo` tinyint(1) NOT NULL DEFAULT '0',
  `apoderado` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `prestamos`
--

CREATE TABLE `prestamos` (
  `id_pre` int(4) NOT NULL,
  `fecha_prestamo` date NOT NULL DEFAULT '0000-00-00',
  `fecha_devolucion` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `id_lib` int(4) NOT NULL,
  `id_alu` int(4) NOT NULL,
  `id_edu` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

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

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuarios`
--

CREATE TABLE `usuarios` (
  `id_usr` int(4) NOT NULL,
  `nombre` varchar(255) CHARACTER SET utf8 COLLATE utf8_spanish2_ci DEFAULT '',
  `correo` varchar(255) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `pass` varchar(32) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `tipo` tinyint(1) NOT NULL DEFAULT '0',
  `cant_agenda` int(4) NOT NULL DEFAULT '0',
  `telefono` varchar(14) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `nmatricula` varchar(20) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `rut` varchar(20) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `apellido1` varchar(100) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `apellido2` varchar(100) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `genero` tinyint(1) NOT NULL DEFAULT '0',
  `reglamento` tinyint(1) NOT NULL DEFAULT '0',
  `fecha_nacimiento` date DEFAULT NULL,
  `fecha_matricula` date DEFAULT NULL,
  `fecha_ingreso` date DEFAULT NULL,
  `direccion` varchar(255) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `fecha_retiro` date DEFAULT NULL,
  `motivo_retiro` tinyint(1) NOT NULL DEFAULT '0',
  `observaciones` text COLLATE utf8_spanish2_ci NOT NULL,
  `telefono2` varchar(14) CHARACTER SET utf8 COLLATE utf8_spanish2_ci NOT NULL DEFAULT '',
  `eliminado` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish2_ci;

--
-- Volcado de datos para la tabla `usuarios`
--

INSERT INTO `usuarios` (`id_usr`, `nombre`, `correo`, `pass`, `tipo`, `cant_agenda`, `telefono`, `nmatricula`, `rut`, `apellido1`, `apellido2`, `genero`, `reglamento`, `fecha_nacimiento`, `fecha_matricula`, `fecha_ingreso`, `direccion`, `fecha_retiro`, `motivo_retiro`, `observaciones`, `telefono2`, `eliminado`) VALUES
(1, 'Eliana Bruzzone', 'elibruzzo@hotmail.com', 'd19ed8f8ac7e5cd3a51a58c3511e6ea4', 0, 0, '', '', '', '', '', 0, 0, '0000-00-00', '0000-00-00', '0000-00-00', '', '0000-00-00', 0, '', '', 0);

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
  MODIFY `id_age` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `cursos`
--
ALTER TABLE `cursos`
  MODIFY `id_cur` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `curso_online`
--
ALTER TABLE `curso_online`
  MODIFY `id_cuo` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `curso_online_items`
--
ALTER TABLE `curso_online_items`
  MODIFY `id_coi` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `libros`
--
ALTER TABLE `libros`
  MODIFY `id_lib` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `prestamos`
--
ALTER TABLE `prestamos`
  MODIFY `id_pre` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `sesiones`
--
ALTER TABLE `sesiones`
  MODIFY `id_ses` int(4) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `usuarios`
--
ALTER TABLE `usuarios`
  MODIFY `id_usr` int(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

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
