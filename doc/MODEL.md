---
html:
  embed_local_images: true
  embed_svg: true
  offline: true
  toc: true
export_on_save:
  html: true
toc:
  depth_from: 2
  depth_to: 10
  ordered: false
print_background: false
---

# Modelo de Datos PADFED

Especificación del modelo de datos de la implementación basada en blockchain de Padrón Federal.

## Contenido

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=10 orderedList=false} -->

<!-- code_chunk_output -->

- [Modelo de Datos PADFED](#modelo-de-datos-padfed)
  - [Contenido](#contenido)
  - [Registro de cambios](#registro-de-cambios)
  - [Convenciones generales](#convenciones-generales)
    - [Campos](#campos)
    - [Formatos](#formatos)
  - [Diagrama general](#diagrama-general)
  - [Objeto: Persona - Persona](#objeto-persona---persona)
    - [Datos comunes](#datos-comunes)
    - [Datos de personas físicas](#datos-de-personas-f%C3%ADsicas)
    - [Datos de personas jurídicas](#datos-de-personas-jur%C3%ADdicas)
    - [Key](#key)
    - [Ejemplos](#ejemplos)
      - [Persona Física {ignore=true}](#persona-f%C3%ADsica-ignoretrue)
      - [Persona Jurídica {ignore=true}](#persona-jur%C3%ADdica-ignoretrue)
  - [Colección: Persona - Impuesto](#colecci%C3%B3n-persona---impuesto)
    - [Ejemplos](#ejemplos-1)
      - [Impuesto Activo (estado AC) {ignore=true}](#impuesto-activo-estado-ac-ignoretrue)
      - [Impuesto con baja definitiva (estado BD) {ignore=true}](#impuesto-con-baja-definitiva-estado-bd-ignoretrue)
  - [Colección: Persona - Domicilio](#colecci%C3%B3n-persona---domicilio)
    - [Key](#key-1)
    - [Ejemplos](#ejemplos-2)
  - [Colección: Persona - Categoría](#colecci%C3%B3n-persona---categor%C3%ADa)
    - [Key](#key-2)
    - [Ejemplos](#ejemplos-3)
  - [Colección: Persona - Contribución Municipal](#colecci%C3%B3n-persona---contribuci%C3%B3n-municipal)
    - [Key](#key-3)
    - [Ejemplos](#ejemplos-4)
  - [Colección: Persona - Actividad](#colecci%C3%B3n-persona---actividad)
    - [Key](#key-4)
    - [Ejemplos](#ejemplos-5)
      - [Actividad primaria (orden 1) {ignore=true}](#actividad-primaria-orden-1-ignoretrue)
      - [Actividad secundaria (orden > 1) {ignore=true}](#actividad-secundaria-orden--1-ignoretrue)
  - [Colección: Persona - Etiqueta](#colecci%C3%B3n-persona---etiqueta)
    - [Key](#key-5)
    - [Ejemplos](#ejemplos-6)
  - [Colección: Persona - Teléfono](#colecci%C3%B3n-persona---tel%C3%A9fono)
    - [Key](#key-6)
  - [Coleccción: Persona - Email](#coleccci%C3%B3n-persona---email)
    - [Key](#key-7)
    - [Ejemplos](#ejemplos-7)
  - [Colección: Persona - Relación](#colecci%C3%B3n-persona---relaci%C3%B3n)
    - [Key](#key-8)
    - [Ejemplos](#ejemplos-8)
  - [Colección: Persona - Jurisdiccion :new:](#colecci%C3%B3n-persona---jurisdiccion-new)
    - [Key](#key-9)
  - [Colección: Persona - Sede Convenio Multilateral :new:](#colecci%C3%B3n-persona---sede-convenio-multilateral-new)
    - [Key](#key-10)
  - [Colección: Persona - Actividad Jurisdiccional :soon:](#colecci%C3%B3n-persona---actividad-jurisdiccional-soon)
    - [Key](#key-11)
  - [Colección: Persona - Domicilio Jurisdiccional :soon:](#colecci%C3%B3n-persona---domicilio-jurisdiccional-soon)
    - [Key](#key-12)
  - [Colección: Persona - Domicilio - Rol :soon:](#colecci%C3%B3n-persona---domicilio---rol-soon)
    - [Key](#key-13)
    - [Ejemplos](#ejemplos-9)
      - [Rol "Fiscal Jurisdiccional" {ignore=true}](#rol-%22fiscal-jurisdiccional%22-ignoretrue)
  - [Colección: Persona - Archivo :soon:](#colecci%C3%B3n-persona---archivo-soon)
  - [Colección: Persona - Fusion :soon:](#colecci%C3%B3n-persona---fusion-soon)
  - [Colección: Persona - Transferencia :soon:](#colecci%C3%B3n-persona---transferencia-soon)
  - [Colección: Persona - Escision :soon:](#colecci%C3%B3n-persona---escision-soon)

<!-- /code_chunk_output -->

## Registro de cambios

| Author         | Date       | Comment                                                    |
| -------------- | ---------- | ---------------------------------------------------------- |
| Pablo Lalloni  | 2019-05-08 | Revisión de máximos, mínimos y ajustes generales, diagrama |
| Fabian Varisco | 2019-05-07 | Versión inicial                                            |

## Convenciones generales

### Campos

- **min** y **max**: Para los strings son longitudes y para los integers son valores.
- **ds**: Fecha de la más reciente modificación del registro en la base de datos de AFIP.
- **org**: Es el código de organismo. `1` es AFIP.

### Formatos

- **#fecha**: Es la representación textual de una fecha con formato `YYYY-MM-DD` y los valores de `DD`, `MM` y `YYYY` deben cumplir las reglas de fechas de calendario estándar.
- Períodos:
  - **#periodomensual**: Formato `YYYYMM`, donde `MM` debe estar en rango [`00`, `12`] e `YYYY` debe estar en el rango [`1000`,`9999`].
  - **#periododiario**: Formato `YYYYMMDD`, donde `MM` debe estar en rango [`00`, `12`] y `DD` puede debe ser `00` si el `MM` es `00` o bien estar en el rango [`01`,`NN`] donde `NN` es la cantidad de días correspondiente al mes `MM` e `YYYY` debe estar en el rango [`1000`,`9999`].

## Diagrama general

![Diagrama](MODEL-DIAGRAM.svg)

## Objeto: Persona - Persona

### Datos comunes

| name     | type     | enum       | min | max | req |
| -------- | -------- | ---------- | --- | --- | --- |
| id       | #cuit    |            |     |     | 🗸  |
| tipoid   | string   | C, E, I, L |     |     | 🗸  |
| tipo     | string   | F, J       |     |     | 🗸  |
| estado   | string   | A, I       |     |     | 🗸  |
| pais     | integer  |            | 100 | 999 |     |
| activoid | #cuit    |            |     |     |     |
| ch       | []string |            |     |     |     |
| ds       | #fecha   |            |     |     |     |

Aclaraciones:

- **activoid**: nueva cuit que se le asignó a la persona
- **ch**: array de nombres de campos cuyos valores fueron modificados en la mas reciente tx

### Datos de personas físicas

| name             | type    | enum | min | max | req |
| ---------------- | ------- | ---- | --- | --- | --- |
| apellido         | string  |      | 1   | 200 | 🗸  |
| nombre           | string  |      | 1   | 200 |     |
| materno          | string  |      | 1   | 200 |     |
| sexo             | string  | M, F |     |     |     |
| documento        | object  |      |     |     |     |
| documento.tipo   | integer |      | 1   | 99  | 🗸  |
| documento.numero | string  |      |     |     | 🗸  |
| nacimiento       | #fecha  |      |     |     |     |
| fallecimiento    | #fecha  |      |     |     |     |

Aclaraciones:

- **materno**: apellido materno

### Datos de personas jurídicas

| name                 | type    | enum | min | max          | req |
| -------------------- | ------- | ---- | --- | ------------ | --- |
| razonsocial          | string  |      | 1   | 200          | 🗸  |
| formajuridica        | integer |      | 1   | 999          |     |
| mescierre            | integer |      | 1   | 12           |     |
| contratosocial       | #fecha  |      |     |              |     |
| duracion             | integer |      | 1   | 999          |     |
| inscripcion          | object  |      |     |              |     |
| inscripcion.registro | integer |      | 1   | 99           |     |
| inscripcion.numero   | integer |      | 1   | 999999999999 | 🗸  |

Aclaraciones:

- **inscripcion**：puede ser en IGJ (registro:1) o en otro registro público de sociedades

### Key

    per:<id>#per

### Ejemplos

#### Persona Física {ignore=true}

Key:

    per:20000000168#per

Objeto:

```json
{
    "id": 20000000168,
    "tipoid": "C",
    "tipo": "F",
    "estado": "A",
    "nombre": "XXXXX",
    "apellido": "XXXXXX",
    "materno": "XXXXX",
    "sexo": "M",
    "nacimiento": "1963-01-01",
    "fallecimiento": "2009-08-02",
    "documento": {
        "tipo": 96,
        "numero": "XX"
    },
    "ds": "2010-02-14"
}
```

#### Persona Jurídica {ignore=true}

Key:

    per:30120013439

Objeto:

```json
{
    "id": 30120013439,
    "tipoid": "C",
    "tipo": "J",
    "estado": "A",
    "razonsocial": "XXXXXXXXX XXXX XX XXXXXXXXXX XXXXXXXXXXX XXXXX",
    "formajuridica": 86,
    "mescierre": 12,
    "contratosocial": "2000-07-31",
    "inscripcion": {
        "registro": 1,
        "numero": 112345
    },
    "ds": "2008-01-21"
}
```

## Colección: Persona - Impuesto

| name             | type            | enum               | min    | max    | req |
| ---------------- | --------------- | ------------------ | ------ | ------ | --- |
| impuesto         | integer         |                    | 1      | 9999   | 🗸  |
| estado           | string          | AC, NA, BD, BP, EX |        |        | 🗸  |
| periodo          | #periodomensual |                    | 100000 | 999912 | 🗸  |
| dia              | integer         |                    | 1      | 31     |     |
| motivo :warning: | #motivo         |                    | 1      | 999999 |     |
| inscripcion      | #fecha          |                    |        |        |     |
| ds               | #fecha          |                    |        |        |     |

Aclaraciones:

- **#motivo**: pendiente cambiar por objeto con estructura como:

```json
{
    "id": "xxxxxx",
    "desde":"2015-02-24"
}
```

Key:

    per:<id>#imp:<impuesto>

### Ejemplos

#### Impuesto Activo (estado AC) {ignore=true}

Key:

    per:20000000168#imp:20

Objeto:

```json
{
    "impuesto": 20,
    "periodo": 200504,
    "estado": "AC",
    "dia": 19,
    "motivo": 44,
    "inscripcion": "2005-04-20",
    "ds": "2015-12-30"
}
```

#### Impuesto con baja definitiva (estado BD) {ignore=true}

Key:

    per:20000000168#imp:5243

Objeto:

```json
{
    "impuesto": 5243,
    "periodo": 201807,
    "estado": "BD",
    "dia": 31,
    "motivo": 557,
    "inscripcion": "2018-06-07",
    "ds": "2018-07-10"
}
```

## Colección: Persona - Domicilio

| name           | type             | enum | min | max     | req |
| -------------- | ---------------- | ---- | --- | ------- | --- |
| tipo           | integer          |      | 1   | 3       | 🗸  |
| orden          | integer          |      | 1   | 9999    | 🗸  |
| estado         | integer          |      | 1   | 99      |     |
| calle          | string           |      |     | 200     |     |
| numero         | integer          |      | 1   | 999999  |     |
| piso           | string           |      |     | 5       |     |
| sector         | string           |      |     | 200     |     |
| manzana        | string           |      |     | 200     |     |
| torre          | string           |      |     | 200     |     |
| unidad         | string           |      |     | 5       |     |
| provincia      | integer          |      | 0   | 24      |     |
| localidad      | string           |      |     | 200     |     |
| cp             | string           |      |     | 8       |     |
| nomenclador    | string :warning: |      |     | 9       |     |
| nombre         | string           |      |     | 200     |     |
| adicional      | object           |      |     |         |     |
| adicional.tipo | integer          |      | 1   | 99      | 🗸  |
| adicional.dato | string           |      |     | 200     | 🗸  |
| baja           | #fecha           |      |     |         |     |
| partido :new:  | integer          |      | 1   | 999     |     |
| partida :new:  | integer          |      | 1   | 9999999 |     |
| ds             | #fecha           |      |     |         |     |

Aclaraciones:

- **unidad** es "Oficina, Departamento o Local"
- **nombre** es "Nombre de Fantasia"
- **partido** es el código del partido provincial
- **partida** es el número de partida inmobiliaria

### Key

    per:<id>#dom:<tipo>.<orden>

### Ejemplos

Key:

    per:20000000168#dom:3.1

Objeto:

```json
{
    "tipo": 3,
    "orden": 1,
    "estado": 1,
    "provincia": 1,
    "localidad": "MAR DEL PLATA SUR",
    "cp": "7600",
    "nomenclador": "1345",
    "calle": "XXXXX",
    "numero": 1000,
    "adicional": {
        "tipo": 3,
        "dato": "XXXXXXX XXXX"
    },
    "nombre": "XX XXXXXX XXXXXX",
    "ds": "2008-01-18"
}
```

## Colección: Persona - Categoría

| name      | type            | enum   | min    | max    | req |
| --------- | --------------- | ------ | ------ | ------ | --- |
| impuesto  | integer         |        | 1      | 9999   | 🗸  |
| categoria | integer         |        | 1      | 999    | 🗸  |
| estado    | string          | AC, BD |        |        | 🗸  |
| periodo   | #periodomensual |        | 100000 | 999912 | 🗸  |
| motivo    | #motivo         |        | 1      | 999999 |     |
| ds        | #fecha          |        |        |        |     |

### Key

    per:<id>#cat:<impuesto>.<categoria>

### Ejemplos

Key:

    per:20000000168#cat:20.1

Objeto:

```json
{
    "impuesto": 20,
    "categoria": 1,
    "periodo": 200004,
    "estado": "AC",
    "ds": "2003-04-14"
}
```

## Colección: Persona - Contribución Municipal

| name      | type    | enum | min | max  | req |
| --------- | ------- | ---- | --- | ---- | --- |
| impuesto  | integer |      | 1   | 9999 | 🗸  |
| municipio | integer |      | 1   | 9999 | 🗸  |
| provincia | integer |      | 0   | 24   | 🗸  |
| desde     | #fecha  |      |     |      | 🗸  |
| hasta     | #fecha  |      |     |      |
| ds        | #fecha  |      |     |      |

### Key

    per:<id>#con:<impuesto>.<municipio>

### Ejemplos

Key:

    per:20000000168#con:5244.98

Objeto:

```json
{
    "impuesto": 5244,
    "municipio": 98,
    "provincia": 3,
    "desde": "2018-06-01",
    "hasta": "2018-07-31",
    "ds": "2018-07-10"
}
```

## Colección: Persona - Actividad

| name           | type    | pattern            | min | max | req |
| -------------- | ------- | ------------------ | --- | --- | --- |
| actividad      | string  | "^883-[0-9]{3,8}$" |     |     | 🗸  |
| orden          | integer |                    | 1   | 999 | 🗸  |
| desde          | #fecha  |                    |     |     | 🗸  |
| hasta          | #fecha  |                    |     |     |     |
| articulo :new: | integer |                    | 1   | 999 |     |
| ds             | #fecha  |                    |     |     |     |

Aclaraciones:

- **actividad**: compuesto por codigo de nomenclador y codigo de actividad

### Key

    per:<id>#act:<actividad>

### Ejemplos

#### Actividad primaria (orden 1) {ignore=true}

Key:

    per:20000000168#act:883-772099

Objeto:

```json
{
    "actividad": "883-772099",
    "orden": 1,
    "desde": 201805,
    "ds": "2018-06-07"
}
```

#### Actividad secundaria (orden > 1) {ignore=true}

Key:

    per:20000000168#act:883-131300

Objeto:

```json
{
    "actividad": "883-131300",
    "orden": 3,
    "desde": 201507,
    "ds": "2015-07-22"
}
```

## Colección: Persona - Etiqueta

Se asimila a la *Caracterización* de AFIP.

| name     | type           | enum   | min      | max      | req |
| -------- | -------------- | ------ | -------- | -------- | --- |
| etiqueta | integer        |        | 1        | 9999     | 🗸  |
| periodo  | #periododiario |        | 10000000 | 99991231 | 🗸  |
| estado   | string         | AC, BD |          |          | 🗸  |
| ds       | #fecha         |        |          |          |     |

### Key

    per:<id>#eti:<etiqueta>

### Ejemplos

Key:

    per:20000000168#eti:160

Objeto:

```json
{
    "etiqueta": 160,
    "periodo": 19940801,
    "estado": "AC",
    "ds": "2003-04-11"
}
```

## Colección: Persona - Teléfono

| name   | type    | enum | min | max             | req |
| ------ | ------- | ---- | --- | --------------- | --- |
| orden  | integer |      | 1   | 999999          | 🗸  |
| pais   | integer |      | 1   | 9999            |     |
| area   | integer |      | 1   | 9999            |     |
| numero | integer |      | 1   | 999999999999999 | 🗸  |
| tipo   | integer |      | 1   | 99              |     |
| linea  | integer |      | 1   | 999             |     |
| ds     | #fecha  |      |     |                 |     |

### Key

    per:<id>#tel:<orden>

Ejemplo:

    per:20000000168#tel:1

```json
{
    "orden": 1,
    "pais": 200,
    "area": 11,
    "numero": 99999999,
    "tipo": 2,
    "linea": 1,
    "ds": "2013-12-16"
}
```

## Coleccción: Persona - Email

| name      | type    | enum | min | max | req |
| --------- | ------- | ---- | --- | --- | --- |
| orden     | integer |      | 1   | 999 | 🗸  |
| direccion | string  |      |     | 100 | 🗸  |
| tipo      | integer |      | 1   | 99  |     |
| estado    | integer |      | 1   | 99  |     |
| ds        | #fecha  |      |     |     |     |

### Key

    per:<id>#ema:<orden>

### Ejemplos

Key:

    per:20000000168#ema:1

Objeto:

```json
{
    "orden": 1,
    "direccion": "XXXXXXXXXXXXXX@XXXXX.XXX.XX",
    "tipo": 1,
    "estado": 2,
    "ds": "2016-10-20"
}
```

## Colección: Persona - Relación

| name    | type    | enum | min | max | req |
| ------- | ------- | ---- | --- | --- | --- |
| persona | #cuit   |      |     |     | 🗸  |
| tipo    | integer |      | 1   | 999 | 🗸  |
| subtipo | integer |      | 1   | 999 | 🗸  |
| desde   | #fecha  |      |     |     | 🗸  |
| ds      | #fecha  |      |     |     |     |

Aclaraciones:

- **tipo**: Inicialmente será siempre `3` que son relaciones societarias.

### Key

    per:<id>#rel:<persona>.<tipo>.<subtipo>

### Ejemplos

Socio de una Sociedad Anónima.

Key:

    per:30120013439#rel:20012531001.3.4

Objeto:

```json
{
    "persona": 20012531001,
    "tipo": 3,
    "subtipo": 4,
    "desde": "2009-01-12",
    "ds": "2014-04-30"
}
```

## Colección: Persona - Jurisdiccion :new:

| name      | type    | enum | min | max | req |
| --------- | ------- | ---- | --- | --- | --- |
| org       | integer |      | 900 | 924 | 🗸  |
| provincia | integer |      | 0   | 24  | 🗸  |
| desde     | #fecha  |      |     |     | 🗸  |
| hasta     | #fecha  |      |     |     |     |
| ds        | #fecha  |      |     |     |     |

### Key

    per:<id>#jur#<org>.<provincia>

Ejemplo:

    per:30120013439#jur:901.0

## Colección: Persona - Sede Convenio Multilateral :new:

| name      | type    | enum | min | max | req |
| --------- | ------- | ---- | --- | --- | --- |
| provincia | integer |      | 0   | 24  | 🗸  |
| desde     | #fecha  |      |     |     | 🗸  |
| hasta     | #fecha  |      |     |     |     |
| ds        | #fecha  |      |     |     |     |

### Key

    per:<id>#cms:<provincia>

Ejemplo:

    per:30120013439#cms:3

## Colección: Persona - Actividad Jurisdiccional :soon:

| name      | type    | pattern                   | min | max | req |
| --------- | ------- | ------------------------- | --- | --- | --- |
| org       | integer |                           | 900 | 924 | 🗸  |
| actividad | string  | `^[0-9]{1,3}-[0-9]{3,8}$` |     |     | 🗸  |
| orden     | integer |                           | 1   | 999 | 🗸  |
| desde     | #fecha  |                           |     |     | 🗸  |
| hasta     | #fecha  |                           |     |     |     |
| articulo  | integer |                           | 1   | 999 |     |
| ds        | #fecha  |                           |     |     |     |

### Key

    per:<id>#acj:<org>.<actividad>

Ejemplo:

    per:20000000168#acj:900.900-12345

## Colección: Persona - Domicilio Jurisdiccional :soon:

| name           | type             | enum | min | max     | req |
| -------------- | ---------------- | ---- | --- | ------- | --- |
| org            | integer          |      | 900 | 924     | 🗸  |
| tipo           | integer          |      | 1   | 3       | 🗸  |
| orden          | integer          |      | 1   | 9999    | 🗸  |
| estado         | integer          |      | 1   | 99      |     |
| calle          | string           |      |     | 200     |     |
| numero         | integer          |      | 1   | 999999  |     |
| piso           | string           |      |     | 5       |     |
| sector         | string           |      |     | 200     |     |
| manzana        | string           |      |     | 200     |     |
| torre          | string           |      |     | 200     |     |
| unidad         | string           |      |     | 5       |     |
| provincia      | integer          |      | 0   | 24      |     |
| localidad      | string           |      |     | 200     |     |
| cp             | string           |      |     | 8       |     |
| nomenclador    | string :warning: |      |     | 9       |     |
| nombre         | string           |      |     | 200     |     |
| adicional      | object           |      |     |         |     |
| adicional.tipo | integer          |      | 1   | 99      | 🗸  |
| adicional.dato | string           |      |     | 200     | 🗸  |
| baja           | #fecha           |      |     |         |     |
| partido :new:  | integer          |      | 1   | 999     |     |
| partida :new:  | integer          |      | 1   | 9999999 |     |
| ds             | #fecha           |      |     |         |     |

Aclaraciones:

- **tipo**: Todos los domicilios jurisdiccionales tendran tipo `3`

Cambios :soon::

- **numero**: cambiaremos a tipo string para permitir adicionar descripciones no numéricas

### Key

    per:<id>#doj:<org>.<tipo>.<orden>

Ejemplo:

    per:20000000168#doj:900.3.20

## Colección: Persona - Domicilio - Rol :soon:

| name            | type    | enum | min | max  | req |
| --------------- | ------- | ---- | --- | ---- | --- |
| org             | integer |      | 1   | 924  | 🗸  |
| tipo            | integer |      | 1   | 3    | 🗸  |
| orden           | integer |      | 1   | 9999 | 🗸  |
| rol             | integer |      | 1   | 99   | 🗸  |
| desde :warning: | #fecha  |      |     |      | 🗸  |
| hasta :warning: | #fecha  |      |     |      |     |
| ds              | #fecha  |      |     |      |     |

### Key

    per:<id>#dor:<org>.<tipo>.<orden>.<rol>

### Ejemplos

#### Rol "Fiscal Jurisdiccional" {ignore=true}

Rol "Fiscal Jurisdiccional" asignado por Córdoba al domicilio con orden 1.

Key:

    per:20000000168#dor:904.1.20.3

## Colección: Persona - Archivo :soon:

Representará registro de archivos documentales almacenados en los sistemas de AFIP.

## Colección: Persona - Fusion :soon:

Represantará datos de fusiones empresarias en las cuales la persona tuvo participación.

## Colección: Persona - Transferencia :soon:

Represantará datos de transferencias de empresas en las cuales la persona tuvo participación.

## Colección: Persona - Escision :soon:

Represantará datos de esciciones empresarias en las cuales la persona tuvo participación.
