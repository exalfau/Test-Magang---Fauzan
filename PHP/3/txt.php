<?php

// Membaca isi file
$filename = 'input.txt';
$fileContent = file_get_contents($filename);

// Menampilkan isi file ke layar
echo $fileContent;

?>
