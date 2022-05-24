<?php

use EverphoneDemo\EverphoneClient;
use EverphoneDemo\EverphoneRandomTextInput;

require_once realpath("vendor/autoload.php");

$client = new EverphoneClient(
    "localhost:8080",
    [
        "credentials" => Grpc\ChannelCredentials::createInsecure()
    ]
);

$call = $client->RandomText((new EverphoneRandomTextInput())->setText(bin2hex(random_bytes(10))));
list($res, $status) = $call->wait();

var_dump($res);
var_dump($status);
