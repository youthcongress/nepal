<?php

$host = '94.136.185.141'; 
$port = '9000';        
$dbname = 'youthcongressnepal';
$user = 'chetanbudathoki';       
$password = 'HeroBudathoki';

try {
    $dsn = "pgsql:host=$host;port=$port;dbname=$dbname;";
    $pdo = new PDO($dsn, $user, $password, [PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
    
    echo "Connected successfully!";
} catch (PDOException $e) {
    echo "Connection failed: " . $e->getMessage();
}

?>
