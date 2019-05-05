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

</style>
