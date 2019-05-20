<template>
    <div id="header-post">
        <v-icon id="menu-icon" href="#"
                v-on:click="active = !active"
                v-bind:class="{ 'menu-active' : active }">fas fa-bars fa-lg</v-icon>

        <v-icon id="menu-icon-tablet" href="#"
                v-on:click="active = !active"
                v-bind:class="{ 'menu-active' : active }">fas fa-bars fa-lg</v-icon>

        <v-icon id="top-icon-tablet" href="#" style="display:none;"
                v-on:click="active = !active"
                v-bind:class="{ 'menu-active' : active }">fas fa-chevron-up fa-lg</v-icon>


        <span id="menu" style="visibility: visible;"
              v-show="active">
            <BaseMenu :items=items></BaseMenu>
            <br>
            <MenuActions
                    :url_previous=actions.url_previous
                    :url_next=actions.url_next>
            </MenuActions>
            <br>

            <SocialSharingList></SocialSharingList>
            <TableOfContents id="toc"
                             :label="tree.label"
                             :nodes="tree.nodes"
                             :href="tree.href"
                             :depth="1">
            </TableOfContents>
        </span>
    </div>
</template>

<script>

import BaseMenu from './BaseMenu.vue';
import MenuActions from './MenuActions.vue';
import SocialSharingList from './SocialSharingList.vue';
import TableOfContents from './TableOfContents.vue';

export default {
    name: 'Navigation',
    components: {
        BaseMenu, MenuActions, SocialSharingList, TableOfContents,
    },
    props: {
        items: Array,
        actions: Object,
    },
    data() {
        return {
            active: false,
            path: window.location.href,
            tree: {
                nodes: [
                    {
                        label: 'item1',
                        href: '/item1',
                        nodes: [
                            {
                                label: 'item1.1',
                                href: '/item1.1',
                            },
                            {
                                label: 'item1.2',
                                href: '/item1.2',
                                nodes: [
                                    {
                                        label: 'item1.2.1',
                                        href: '/item1.2.1',
                                    },
                                ],
                            },
                        ],
                    },
                    {
                        label: 'item2',
                        href: '/item2',
                    },
                ],
            },
        };
    },
};
</script>


<style>
    #header-post #menu-icon:hover {
        color: #2bbc8a;
    }

    #header-post #menu-icon-tablet:hover {
        color: #2bbc8a;
    }
    #header-post #top-icon-tablet {
        position: fixed;
        right: 2rem;
        bottom: 2rem;
        margin-right: 2rem;
        margin-left: 15px;
    }
    #header-post #top-icon-tablet:hover {
        color: #2bbc8a;
    }
    #header-post .active {
        color: #2bbc8a;
    }

    #menu-icon, #menu-icon-tablet, #top-icon-tablet {
        margin-top: 5px;
    }

    .menu-inactive {
        color: #c9cacc;
    }
    .menu-active {
        color: #2bbc8a;
    }

    #toc {
        overflow: auto;
        margin-top: 1rem;
        padding-right: 2rem;
        max-width: 20em;
        max-height: calc(95vh - 7rem);
        text-align: right;
    }
</style>
