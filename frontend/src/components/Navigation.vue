<template>
    <div id="header-post">
        <a id="menu-icon" href="#" class="active"><i class="fas fa-bars fa-lg"></i></a>
        <a id="menu-icon-tablet" href="#" class="active"><i class="fas fa-bars fa-lg"></i></a>
        <a id="top-icon-tablet" href="#" style="display:none;"><i class="fas fa-chevron-up fa-lg"></i></a>
        <span id="menu" style="visibility: visible;">
            <BaseMenu :items=items></BaseMenu>
            <br>
            <MenuActions
                    :url_previous=actions.url_previous
                    :url_next=actions.url_next>
            </MenuActions>
        <br>
        <div id="share" style="display: none">
          <ul>
    <!--  <li><a class="icon" href=""><i class="fab fa-facebook " aria-hidden="true"></i></a></li>  social sharing here-->
      </ul>

        </div>
      </span>
    </div>
</template>

<script>

import BaseMenu from './BaseMenu.vue';
import MenuActions from './MenuActions.vue';

export default {
    name: 'Navigation',
    components: { BaseMenu, MenuActions },
    props: {
        items: Array,
        actions: Object,
    },
};
</script>
