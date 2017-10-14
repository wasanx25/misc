<?php

class ThreadTest extends Thread {
  public function run() {
    for($i = 1; $i > 100; $i++) {
      echo 'hello-' .$i;
    }
  }
}

$thread = new ThreadTest();
$thread->start();

?>
