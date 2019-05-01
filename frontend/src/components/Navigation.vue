<template>
    <div id="header-post">
        <a id="menu-icon" href="#"><i class="fas fa-bars fa-lg"></i></a>
        <a id="menu-icon-tablet" href="#"><i class="fas fa-bars fa-lg"></i></a>
        <a id="top-icon-tablet" href="#" style="display:none;"><i class="fas fa-chevron-up fa-lg"></i></a>
        <span id="menu">
    <span id="nav">
      <ul>

          <li><a href="...">Home</a></li>
          <li><a href="...">Writing</a></li>
          <li><a href="...">About</a></li>

      </ul>
    </span>
    <br>
    <span id="actions">
      <ul>

        <li><a class="icon" href=""><i class="fas fa-chevron-right" aria-hidden="true" ></i></a></li>
        <li><a class="icon" href="#" ><i class="fas fa-chevron-up" aria-hidden="true"></i></a></li>
        <li><a class="icon" href="#"><i class="fas fa-share-alt" aria-hidden="true"></i></a></li>
      </ul>
      <span id="i-prev" class="info" style="display:none;">Previous post</span>
      <span id="i-next" class="info" style="display:none;">Next post</span>
      <span id="i-top" class="info" style="display:none;">Back to top</span>
      <span id="i-share" class="info" style="display:none;">Share post</span>
    </span>
    <br>
    <div id="share" style="display: none">
      <ul>
        <li><a class="icon" href=""><i class="fab fa-facebook " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-twitter " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-linkedin " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-pinterest " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fas fa-envelope " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-get-pocket " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-reddit " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-stumbleupon " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-digg " aria-hidden="true"></i></a></li>
        <li><a class="icon" href=""><i class="fab fa-tumblr " aria-hidden="true"></i></a></li>
      </ul>

    </div>
    <div id="toc">
      <li class="toc-item toc-level-1"><a class="toc-link" href="#Level-1">
          <span class="toc-number">Level 1</span>
          <span class="toc-text">Text</span></a>
        <ol class="toc-child"><li class="toc-item toc-level-2"><a class="toc-link" href="#Level-2">
              <span class="toc-number">1.</span>
              <span class="toc-text">Level 2</span></a>
            <ol class="toc-child"><li class="toc-item toc-level-3"><a class="toc-link" href="#Level-3">
                  <span class="toc-number">1.1.</span>
                  <span class="toc-text">Level 3</span></a>
              </li>
            </ol></li>
        </ol>
      </li>
    </div>
  </span>
</div>

</template>

<script>

export default {
    name: 'Navigation',
};
</script>
