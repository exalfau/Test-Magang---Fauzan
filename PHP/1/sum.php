<?php

function getSum($arr) {
    $sum = 0;
    foreach ($arr as $num) {
        $sum += $num;
    }
    return $sum;
}

// Contoh penggunaan fungsi
$array = [1, 2, 3, 4, 5];
$hasil = getSum($array);
echo "Hasil penjumlahan: " . $hasil;

?>
