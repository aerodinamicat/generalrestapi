# General REST API

## Mi proyecto personal sin rumbo ni destino

### Propósito general
Para el momento que escribo éste *README.md* aún no tengo claro para nada que funciones
o propósito general tendrá todo este contenido; pero si puedo afirmar y confirmar que
se realizará completamente en lenguaje *Go*.


### Tareas
Por el momento, comenzaré por crear una REST API (sin identificación y sin seguridad):
- CRUD para objetos:
    - *person* (Persona)
    - *user* (Usuario)
    - *ledger account* (Cuenta contable)


### Arquitectura usada y entorno de trabajo
Mi arquitectura:
- Sistema operativo: Ubuntu 22.04 LTS

Mi *workspace* (entorno de trabajo) actualmente se compone por:
- Visual Studio Code

Como se puede ver no tengo muy claro qué compone ni mi *arquitectura* ni mi *entorno de
trabajo*.


### Diario
Y por llevar el registro de un diario:
2022-06-29:1
- Crearé la carpeta padre del proyecto local, llamada *generalrestapi* con:
    > mkdir generalrestapi
- Crearé este mismo archivo, he incluiré éste contenido en notación *Mark down*, con:
    > touch README.md
- Crearé éste repositorio en *GitHub*, con el nombre de *generalrestapi*.
- Iniciaré el módulo de *Go* con:
    > go mod init github.com/aerodinamicat/generalrestapi
- Iniciaré el repositorio de git localmente con:
    > git init
- Añadiré todos los archivos contenidos en la carpeta local del proyecto a git con:
    > git add .
- Realizaré el primer commit con:
    > git commit -m "Initial commit - Full push"
- Añadiré localmente este mismo repositorio en remoto, teniendo previamente configurado
el acceso a *GitHub* con SSH, con:
    > git remote add origin git@github.com:aerodinamicat/generalrestapi.git
- Y, por último en primera instancia, realizaré el primer *push* al recién creado
repositorio con:
    > git push origin main

2022-06-29:1
- Creado directorio de modelos *models* con:
    > mkdir models
- Creado archivo de modelo *person.go* en el directorio de modelos con:
    > touch models/person.go
- Implementado el modelo *person* y comentado el código escrito.
- Actualizado *README.md* y cambiado los títulos *h3* y *h5* a *h2* y *h3* respectivamente.

### Conclusiones
Este proyecto constituye un lugar donde comenzar a realizar mis primeros pasos, no en
el mundo de la programación, pero si en las nuevas tecnologías y en el lenguaje *Go*.