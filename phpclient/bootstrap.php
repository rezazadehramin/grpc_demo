<?php

use EverphoneDemo\EverphoneClient;
use EverphoneDemo\EverphoneRandomTextInput;

require_once realpath("vendor/autoload.php");

$client = new EverphoneClient(
    "goserver:8080",
    [
        "credentials" => Grpc\ChannelCredentials::createInsecure()
    ]
);

while (true) {
    $value = bin2hex(random_bytes(10));
    $call = $client->RandomText((new EverphoneRandomTextInput())->setText($value));
    list($reply, $status) = $call->wait();

    echo $reply->getText() . "\n";
    sleep(3);
}
