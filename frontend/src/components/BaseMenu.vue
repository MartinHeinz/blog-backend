<template>
    <span id="nav">
        <ul>
            <BaseMenuItem v-for="item in items"
                          v-bind:url="item.url"
                          v-bind:value="item.value"
                          v-bind:key="item.id">
            </BaseMenuItem>
        </ul>

    </span>
</template>

<script>

import BaseMenuItem from './BaseMenuItem.vue';

export default {
    name: 'BaseMenu',
    components: { BaseMenuItem },
    props: {
        items: Array,
    },
};
</script>

<style>
    ul {
        display: inline-block;
        margin: 0;
        padding: 0;
        list-style-type: none;
    }

    ul li {
        display: inline-block;
        margin-right: 15px;
        vertical-align: middle;
        border-right: 1px dotted #2bbc8a;
        text-align: -webkit-match-parent;
    }

    li:last-child {
        margin-right: 0;
        border-right: 0;
    }
    #header-post #nav ul li:last-child a {
        margin-right: 0;
    }

    ul a:hover {
        background-image: linear-gradient(transparent, transparent 5px, #2bbc8a 5px, #2bbc8a);
        background-position: bottom;
        background-size: 100% 6px;
        background-repeat: repeat-x;
    }

    ul li:last-child a {
        margin-right: 0;
    }

    ul {
        line-height: 15px;
    }
    ul a {
        margin-right: 15px;
        color: #2bbc8a;
        text-decoration: none;
    }
</style>
