package rabbit

import (
        "os"
        "github.com/streadway/amqp"
        "errors"
        "github.com/Zombispormedio/smart-push/lib/utils"
)


type Rabbit struct{
    Conn *amqp.Connection
    
    Chan *amqp.Channel
    
    ExName string
    ExType string
    
    
}

func New(exname string, extype string) (*Rabbit, error){
    var Error error
    rabbit:= Rabbit{}
    rabbit.Conn, Error=amqp.Dial(os.Getenv("RABBIT"));
    
    if Error != nil{
        return  nil, Error
    }
    
    rabbit.Chan, Error= rabbit.Conn.Channel()
    
    if Error != nil{
        return  nil, Error
    }
    
    Error=rabbit.Exchange(exname, extype)
   
    
    return rabbit, Error 
}

func (rabbit *Rabbit) Exchange(exname string, extype string) error{
    var Error error
    
    if exname == nil{
        return errors.New("You Need Exname")
    }
    
    if extype== nil{
          return errors.New("You Need ExType")
    }
    
    
    
    return Error
    
}

func (rabbit *Rabbit) Close(){
    rabbit.Chan.Close();
    rabbit.Conn.Close();
}