# General REST API
Mi proyecto personal sin rumbo ni destino

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


### Conclusiones
Este proyecto constituye un lugar donde comenzar a realizar mis primeros pasos, no en
el mundo de la programación, pero si en las nuevas tecnologías y en el lenguaje *Go*.


### Ideas y tareas pendientes posibles/probables
Ésta sección se concentrará en dejar escritas ideas a discutir, así como tareas que
podrían ser satisfechas:<ul>
    <li>[] Creación e implementación del archivo *.gitignore*</li>
    <li>[] Pruebas unitarias sobre el modelo de *person.go*.</li>
</ul>


### Diario
Y por llevar registro de un diario:

2022-06-29:4
- Actualizado *README.md*:
    - Añadida la sección 'Ideas y tareas pendientes posibles/probables'.
    - Intento de dar estilo con cajas de confirmación tipo *check box*.
- Para el desarrollo de éste proyecto, he pensado usar un contenedor de docker con una
imagen de PostGreSQL para el servicio de base de datos. Creado el archivo de docker
*Dockerfile* en el directorio de base de datos con:
    > touch database/Dockerfile
- Ya que he pensado en usar diferentes servicios, he creado el archivo docker-compose
en el directorio principal del proyecto *docker-compose.yml* con
    > touch docker-compose.yml
Aunque por ahora solo existe, y se vaya a utilizar, únicamente el servicio de base de
datos, no está mal ir realizando éste tipo de tareas con vistas al futuro.
- Mediante el archivo de docker-compose, ubicado en el directorio raíz del proyecto, he
realizado pruebas de lanzamiento y conexión al servicio dockerizado de postgres, con:
    > docker-compose build && docker-compose up
Y para parar el contenedor lanzado, así como el resto de componentes, además de borrar
las imágenes de compilación creadas, con:
    > docker-compose down --rmi local
- Sometido a pruebas:
    - [x] Lanzamiento de base de datos postgres dockerizada.

2022-06-29:3
- Actualizado *README.md* y cambiado:
    - Subtítulo principal de *h2* a texto normal.
    - Adelantada sección 'Conclusiones' justo antes de sección 'Diario'
    - Cambiado el orden de entradas de diario de arriba a abajo, de mas reciente a mas
    antigüo.
- Creado el directorio de base de datos *database* con:
    > mkdir database
- Creado el archivo de consulta de base de datos inicial *initialQuery.sql* en el 
directorio de base de datos con:
    > touch database/initialQuery.sql
- Implementado el modelo, o tabla, de base de datos 'persons' y comentado el código
escrito.
- Cambiada la formación del mensaje de los *commit*, de ahora en adelante se nombrarán
como las entradas de éste diario, salvo excepciones puntuales. (Ergo este *commit* será
'2022-06-29:3')

2022-06-29:2
- Creado el directorio de modelos *models* con:
    > mkdir models
- Creado el archivo de modelo *person.go* en el directorio de modelos con:
    > touch models/person.go
- Implementado el modelo de *Go* *person* y comentado el código escrito.
- Actualizado *README.md* y cambiado los títulos *h3* y *h5* a *h2* y *h3*
respectivamente.

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
- Añadiré localmente este mismo repositorio en remoto, teniendo previamente
configurado el acceso a *GitHub* con SSH, con:
    > git remote add origin git@github.com:aerodinamicat/generalrestapi.git
- Y, por último en primera instancia, realizaré el primer *push* al recién creado
repositorio con:
    > git push origin main