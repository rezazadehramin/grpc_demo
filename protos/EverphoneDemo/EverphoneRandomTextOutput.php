<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: everphone.proto

namespace EverphoneDemo;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>EverphoneRandomTextOutput</code>
 */
class EverphoneRandomTextOutput extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string Text = 1;</code>
     */
    protected $Text = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $Text
     * }
     */
    public function __construct($data = NULL) {
        \EverphoneDemo\Meta\Everphone::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string Text = 1;</code>
     * @return string
     */
    public function getText()
    {
        return $this->Text;
    }

    /**
     * Generated from protobuf field <code>string Text = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setText($var)
    {
        GPBUtil::checkString($var, True);
        $this->Text = $var;

        return $this;
    }

}

