<?php

function reverseWords($str) {
    // Memisahkan kata-kata dalam string menjadi array
    $words = explode(' ', $str);
    
    // Membalikkan urutan kata dalam array
    $reversedWords = array_reverse($words);
    
    // Menggabungkan kembali kata-kata menjadi string
    $reversedStr = implode(' ', $reversedWords);
    
    return $reversedStr;
}

// Contoh penggunaan fungsi
$string = "Perkenalkan nama saya Fauzan!.";
$hasil = reverseWords($string);
echo $hasil;

?>
