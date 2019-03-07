<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: tabletmanagerdata.proto

namespace Vitess\Proto\Tabletmanagerdata {

  class ExecuteHookResponse extends \DrSlump\Protobuf\Message {

    /**  @var int */
    public $exit_status = null;
    
    /**  @var string */
    public $stdout = null;
    
    /**  @var string */
    public $stderr = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'tabletmanagerdata.ExecuteHookResponse');

      // OPTIONAL INT64 exit_status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "exit_status";
      $f->type      = \DrSlump\Protobuf::TYPE_INT64;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING stdout = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "stdout";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING stderr = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "stderr";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <exit_status> has a value
     *
     * @return boolean
     */
    public function hasExitStatus(){
      return $this->_has(1);
    }
    
    /**
     * Clear <exit_status> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteHookResponse
     */
    public function clearExitStatus(){
      return $this->_clear(1);
    }
    
    /**
     * Get <exit_status> value
     *
     * @return int
     */
    public function getExitStatus(){
      return $this->_get(1);
    }
    
    /**
     * Set <exit_status> value
     *
     * @param int $value
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteHookResponse
     */
    public function setExitStatus( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <stdout> has a value
     *
     * @return boolean
     */
    public function hasStdout(){
      return $this->_has(2);
    }
    
    /**
     * Clear <stdout> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteHookResponse
     */
    public function clearStdout(){
      return $this->_clear(2);
    }
    
    /**
     * Get <stdout> value
     *
     * @return string
     */
    public function getStdout(){
      return $this->_get(2);
    }
    
    /**
     * Set <stdout> value
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteHookResponse
     */
    public function setStdout( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <stderr> has a value
     *
     * @return boolean
     */
    public function hasStderr(){
      return $this->_has(3);
    }
    
    /**
     * Clear <stderr> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteHookResponse
     */
    public function clearStderr(){
      return $this->_clear(3);
    }
    
    /**
     * Get <stderr> value
     *
     * @return string
     */
    public function getStderr(){
      return $this->_get(3);
    }
    
    /**
     * Set <stderr> value
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\ExecuteHookResponse
     */
    public function setStderr( $value){
      return $this->_set(3, $value);
    }
  }
}

