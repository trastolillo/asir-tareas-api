create table if not exists Modulos (
  IdModulo varchar(6) primary key not null,
  Modulo varchar(200) not null,
  URL varchar(250) not null
) engine = InnoDB;
create table if not exists Tareas (
  IdTarea int not null auto_increment,
  IdModulo varchar(50) not null,
  URL varchar(255),
  Unidad tinyint not null,
  Tipo enum('cuestionario', 'tarea') not null,
  FechaLimite date not null,
  FechaTerminado date,
  primary key (IdTarea),
  constraint fk_modulo foreign key(IdModulo) references Modulos(IdModulo) on delete restrict on update cascade
) engine = InnoDB;