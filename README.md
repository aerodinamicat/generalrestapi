# General REST API
Mi proyecto personal sin rumbo ni destino

### Propósito general
Para el momento que escribo éste *README.md* aún no tengo claro para nada que funciones
o propósito general tendrá todo este contenido; pero si puedo afirmar y confirmar que
se realizará completamente en lenguaje *Go*.


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


### Procedimiento de ejecución del proyecto
Esta sección tiene como objetivo proveer una serie de comandos para obtener el proyecto
en un ámbito local, así como realizar la prueba de sus funcionalidades disponibles hasta
su momento presente. PRÓXIMAMENTE


### Tareas pendientes a realizar
Por el momento, comenzaré por crear una REST API (sin identificación y sin seguridad):
- C.R.U.D. para objetos:
    - *person* (Persona)
        - [x] Definición e implementación del objeto en *Go*.
        - [x] Pruebas unitarias sobre el objeto de *Go*.
        - [x] Creada Sentencia SQL.
        - [ ] List (usando el método de HTTP: 'GET').
        - [ ] Get (usando el método de HTTP: 'GET').
        - [ ] Create (usando el método de HTTP: 'POST').
        - [ ] Update (usando el método de HTTP: 'PUT').
        - [ ] Delete (usando el método de HTTP: 'DELETE').
    - *user* (Usuario)
        - [x] Definición e implementación del objeto en *Go*.
        - [x] Creada Sentencia SQL.
        - [ ] List (usando el método de HTTP: 'GET').
        - [ ] Get (usando el método de HTTP: 'GET').
        - [ ] Create (usando el método de HTTP: 'POST').
        - [ ] Update (usando el método de HTTP: 'PUT').
        - [ ] Delete (usando el método de HTTP: 'DELETE').
    - *ledger account* (Cuenta contable)
        - [ ] Definición e implementación del objeto en *Go*.
        - [ ] Creada Sentencia SQL.
        - [ ] List (usando el método de HTTP: 'GET').
        - [ ] Get (usando el método de HTTP: 'GET').
        - [ ] Create (usando el método de HTTP: 'POST').
        - [ ] Update (usando el método de HTTP: 'PUT').
        - [ ] Delete (usando el método de HTTP: 'DELETE').
    - *company* (Empresa)
        - [ ] Definición e implementación del objeto en *Go*.
        - [x] Creada Sentencia SQL.
        - [ ] List (usando el método de HTTP: 'GET').
        - [ ] Get (usando el método de HTTP: 'GET').
        - [ ] Create (usando el método de HTTP: 'POST').
        - [ ] Update (usando el método de HTTP: 'PUT').
        - [ ] Delete (usando el método de HTTP: 'DELETE').


### Ideas posibles y/o probables
Ésta sección se concentrará en dejar escritas ideas a discutir, así como tareas que
podrían ser satisfechas:
    - [ ] Creación e implementación del archivo *.gitignore*.
    - [ ] Completar sección 'Procedimiento de ejecución del proyecto'


### Diario
Y por llevar registro de un diario:

2022-07-01:3
- Actualizado *README.md*:
    - Corregidos algunos *typos* dispersos por este documento.
    - Cambiado sección 'Tareas pendientes a realizar - C.R.U.D para objetos', todas
    las referencias a métodos 'PATCH' de *HTTP* por 'PUT'.
    - Removidos los caracteres de escape '`' en los enlaces de éste mismo archivo.
- Implementado el modelo *PersonsServer* de *Go* para satisfacer el servicio
*PersonService* ubicado en el archivo *pbModels.go*.
- Creado el archivo *databaseRepository.go* de *Go* en el directorio correspondiente
con:
    > `$ touch repositories/databaseRepository.go`
- Implementado el modelo *DatabaseRepository* de *Go*. No se ha realizado prueba alguna.
- Creado el archivo *postgres.go* de *Go* en el directorio 'database' con:
    > `$ touch database/postgres.go`
- Importadas algunas librerías externas:
    - Para trabajar con *PostgreSQL* mediante *Go* con:
        > `$ go get github.com/lib/pq`
    - Para crear cadenas de caracteres únicos firmados mediante *Go* con:
        > `$ go get github.com/segmentio/ksuid`
    - Para manejar las rutas y crear los diferentes 'endpoints' mediante *Go* con:
        > `$ go get github.com/gorilla/mux`
    - Para transmitir de forma compacta, autónoma y segura información *JSON* con:
        > `$ go get github.com/golang-jwt/jwt`
    - Para trabajar con *WebSockets* mediante *Go* con:
        > `$ go get github.com/gorilla/websocket`

2022-07-01:2
- Modificaciones en el archivo *pbModels.proto* los siguientes errores:
    - Corregidos algunos *typos* de referencia de métodos.
    - Cambiado el nombre del paquete de *Protobuffer* a:
        - 'pbmodels'.
    - Cambiado el nombre del paquete de *Go* a:
        - 'github.com/aerodinamicat/generalrestapi/pbmodels'
- Importadas algunas librerías externas para comenzar a trabajar con *Protobuffer* con:
    > `$ go get google.golang.org/grpc`
    >
    > `$ go get google.golang.org/protobuf`
- Realizada la compilación, por primera vez, del archivo del modelos *pbModels.go* de
*Protobuffer* con (desde la carpeta 'raíz' del proyecto):
    > `$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative models-proto/pbModels.go`
- Creado el directorio *repositories* en el directorio 'raíz' del proyecto con:
    > `$ mkdir repositories`
- Creado el archivo *personsRepository.go* de *Go* en el directorio correspondiente con:
    > `$ touch repositories/personServiceRepository.go`
- Implementado el modelo *PersonService* de *Protobuffer* en el archivo
*personRepository.go* de *Go* para satisfacer todos los métodos definidos en el archivo
*pbModels.proto*.
- Creado el directorio *servers* en el directorio 'raíz' del proyecto con:
    > `$ mkdir servers`
- Creado el archivo *PersonsServer.go* de *Go* en el directorio correspondiente con:
    > `$ touch servers/personServer.go`


2022-07-01:1
- Actualizado *README.md*.
    - Actualizada sección 'Tareas pendientes a realizar' tras leer:
        -  [Métodos estándar - API de Google Cloud](https://cloud.google.com/apis/design/standard_methods)
    - Corregidas, y estandarizadas, las formulaciones de referencias, escritas en este
    archivo, sobre:
        - Entidades de *SQL*.
        - Modelos de *Go*.
        - Ramas de *Git*.
        - Archivos de *Docker*
        - Archivos de *Docker Compose*
- Creado el directorio *models-protos* en el directorio 'raíz' del proyecto con:
    > `$ mkdir models-protos`
- Creado el archivo de modelo *pbModels.proto* de *Protobuffer* en su respectivo
directorio con:
    > `$ touch models-protos/pbModels.proto`
- Implementado el modelo 'Person' de *Protobuffer* en el archivo *pbModels.proto*. No
ha sido compilado aún por primera vez.


2022-06-30:3
- Actualizado *README.md*.
    - Corregidos algunos *typos* sueltos.
- Modificado el modelo *AccountingTransaction* de *Go*:
    - Eliminada la propiedad *ZeroSum*, ya que representaba un valor de un campo calculado.
    No tiene sentido guardar su valor.
- Modificado el archivo *initialQuery.sql* de *SQL* para añadir las entidades de:
    - *users*
    - *logbook_records*
    - *companies*
    - *accounting_records*
    - *accounting_transactions*
- Una vez añadida la entidad 'logbook_records' de *SQL*, han sido eliminados del resto
de entidades (ya que es función de 'logbook_records' guardar dicha información) los campos
de:
    - *created_at*
    - *created_by_user_id*
    - *modified_at*
    - *modified_by_user_id*
- Modificado el modelo *AccountingRecord* de *Go*:
    - Renombrada la propiedad *Type* a *Nature*.
    - Renombradas las constantes de ámbito global:
        - *TYPE_CREDIT*, renombrado a: *NATURE_CREDIT*.
        - *TYPE_DEBIT*, renombrado a: *NATURE_DEBIT*.


2022-06-30:2
- Actualizado *README.md*: Corregidos algunos *typos* sueltos.


2022-06-30:1
- Actualizado *README.md*:
    - Corregidos algunos *typos* sueltos.
    - Añadida la sección 'Procedimiento de ejecución del proyecto'
- Creado el archivo de modelo *logbook.go* de *Go*, cuyo objetivo es conservar un
histórico de 'cuando' y 'quién' realizó cambios en algún objeto registrado en la 'DB',
con:
    > `$ touch models/logbook.go`
- Implementado el modelo de *Logbook* de *Go*. No se han realizado pruebas dado que
no tiene métodos, o comportamientos. Modificados el resto de los modelos, previamente
creados, para que usen éste nuevo.
- Creado el archivo de modelo *user.go* de *Go* con:
    > `$ touch models/user.go`
- Implementado el modelo de *User* de *Go*. No se han realizado pruebas dado que éste
modelo en particular carece de métodos, o comportamientos.
- Creado el archivo de modelo *accountingRecord.go* de *Go* en su directorio correspondiente
con:
    > `$ touch models/accountingRecord.go`
- Implementado el modelo de *AccountingRecord*. No se han realizado pruebas dado que
éste modelo en particular carece de métodos, o comportamientos.
- Creado el archivo de modelo *accountingTransaction.go* de *Go* en su directorio
correspondiente con:
    > `$ touch models/accountingTransaction.go`
- Implementado el modelo de *AccountingTransaction* de *Go*. Se han realizado pruebas
satisfactorias a su único método (por el momento a fecha de hoy):
    - *CalculateZeroSum(aRecords)*


2022-06-29:6
- Actualizado *README.md*:
    - Corregido algunos errores de salto de línea en las *blockquotes*.
- Creado el directorio de pruebas *tests* en el directorio 'raíz' del proyecto con:
    > `$ mkdir tests`
- Creado el archivo de tests *person_test.go* de *Go* en su directorio correspondiente
para la realización de pruebas sobre el modelo *Person* de *Go* con:
    > `$ touch tests/person_test.go`
- Implementadas funciones de prueba en el archivo de tests *person_test.go* de *Go* para
los métodos, o comportamientos, del modelo *Person* de *Go*:
    - *GetFullName()*
    - *GetFullNameToList()*
    - *GetInitials(bool)*
    - *GetAge()*
- Ejecutadas, y con éxito, una batería de pruebas para los métodos anteriores en el
modelo *Person* de *Go* con:
    > `$ go test tests/person_test.go`
- Se creó la rama *person* en *git* porque durante la creación e implementación de sus
tests de pruebas se encontraron algunas dificultades. Ahora que todo se ha solventado
satisfactoriamente, se ha realizado un *merge* a la rama *main* con:
    > `$ git merge main`


2022-06-29:5
- Actualizado *README.md*:
    - Corregido intento de estilo de cajas de confirmación tipo *check box* en la
    sección 'Ideas posibles y/o probables. Se ha reescrito en *Mark down*, removiendo
    el código *HTML* previamente implementado.
    - Modificación de sección 'Tareas pendientes para realizar':
        - Añadido estilo de cajas de confirmación.
        - Añadidas algunas tareas a realizar
    - Modificadas las *Blockquotes*. Añadida apariencia de bloque de código.


2022-06-29:4
- Actualizado *README.md*:
    - Añadida la sección 'Ideas y tareas pendientes posibles/probables'.
    - Intento de dar estilo con cajas de confirmación tipo *check box*.
- Para el desarrollo de éste proyecto, he pensado usar un contenedor de docker con una
imagen de *PostGreSQL* para el servicio de base de datos. Creado el archivo *Dockerfile*
de *Docker* en el directorio de base de datos con:
    > `$ touch database/Dockerfile`
- Ya que he pensado en usar diferentes servicios, he creado el archivo *docker-compose.yml*
de *Docker-Compose* en el directorio principal del proyecto con:
    > `$ touch docker-compose.yml`
- Aunque por ahora solo existe, y se vaya a utilizar, únicamente el servicio de base de
datos, no está mal ir trabajando y realizando éste tipo de procedimientos y tareas con
vistas al futuro.
- Mediante el archivo de *docker-compose.yml* de *Docker-Compose*, ubicado en el directorio
raíz del proyecto, he realizado pruebas de lanzamiento y conexión, mediante *DBeaver* al
servicio dockerizado de *PostgreSQL*, con:
    > `$ docker-compose build && docker-compose up`
- Y para detener el contenedor lanzado, así como el resto de componentes, además de borrar
las imágenes de compilación creadas, con:
    > `$ docker-compose down --rmi local`
- Sometido a pruebas:
    - [x] Lanzamiento de base de datos de *PostgreSQL* dockerizada.
    - [x] Conexión mediante *DBeaver*, un RDBMS (Sistema de Gestión de Base de Datos
    Relacional), realizado de forma satisfactoria.


2022-06-29:3
- Actualizado *README.md* y cambiado:
    - Subtítulo principal de *h2* a texto normal.
    - Adelantada sección 'Conclusiones' justo antes de sección 'Diario'
    - Cambiado el orden de entradas de diario de arriba a abajo, de mas reciente a mas
    antigüo.
- Creado el directorio de base de datos *database* en el directorio 'raíz' del proyecto
con:
    > `$ mkdir database`
- Creado el archivo de consulta de base de datos *initialQuery.sql* de *SQL* en su directorio
correspondiente con:
    > `$ touch database/initialQuery.sql`
- Implementada la entidad , o tabla, de base de datos 'persons' de *SQL* y comentado el
código escrito.
- Cambiada la formación del mensaje de los *commit*. De ahora en adelante se nombrarán
como las entradas de éste diario, salvo excepciones puntuales. (Ergo este *commit* será
'2022-06-29:3')

2022-06-29:2
- Actualizado *README.md*: cambiado los títulos *h3* y *h5* a *h2* y *h3* respectivamente.
- Creado el directorio de modelos *models* en el directorio 'raíz' del proyecto, para los
modelos de *Go* con:
    > `$ mkdir models`
- Creado el archivo *person.go* de modelo de *Go* en su directorio correspondiente con:
    > `$ touch models/person.go`
- Implementado el modelo *Person* de *Go* y comentado el código escrito.


2022-06-29:1
- Creado el directorio *generalrestapi* que será el directorio 'raíz' del proyecto local
con:
    > `$ mkdir generalrestapi`
- Creado el archivo *README.md* (este mismo), he incluiré éste contenido, en notación
*Mark down*, con:
    > `$ touch README.md`
- Creado el repositorio *generalrestapi* de *GitHub* mediante su plataforma web.
- Iniciado el módulo *github.com/aerodinamicat/generalrestapi* de *Go* con:
    > `$ go mod init github.com/aerodinamicat/generalrestapi`
- Iniciado el repositorio de git localmente con:
    > `$ git init`
- Añadidos todos los archivos contenidos en la carpeta local del proyecto a git con:
    > `$ git add .`
- Realizado el primer commit con:
    > `$ git commit -m "Initial commit - Full push"`
- Añadido el repositorio remoto *origin* de *Git*, teniendo previamente configurado
el acceso a *GitHub* con SSH, con:
    > `$ git remote add origin git@github.com:aerodinamicat/generalrestapi.git`
- Y, por último, realizado el primer *push* al recién creado repositorio de *GitHub* con:
    > `$ git push origin main`
