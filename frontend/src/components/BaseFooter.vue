<template>
    <footer id="footer">
        <div class="footer-left">
            Copyright &copy; 2019 Martin Heinz
        </div>
        <div class="footer-right">
            <nav>
                <ul>
                    <BaseFooterItem v-for="item in items"
                                    v-bind:url="item.url"
                                    v-bind:value="item.value"
                                    v-bind:key="item.id">
                    </BaseFooterItem>
                </ul>
            </nav>
        </div>
    </footer>
</template>

<script>

import BaseFooterItem from './BaseFooterItem.vue';

export default {
    name: 'BaseFooter',
    components: { BaseFooterItem },
    props: {
        items: Array,
    },
};
</script>

<style>
    #footer {
        position: absolute;
        bottom: 0;
        margin-bottom: 10px;
        width: 100%;
        color: #666;
        vertical-align: top;
        text-align: center;
        font-size: 11px;
    }
    #footer ul {
        margin: 0;
        padding: 0;
        list-style: none;
    }
    #footer li {
        display: inline-block;
        margin-right: 15px;
        border-right: 1px solid #666;
        vertical-align: middle;
    }
    #footer li a {
        margin-right: 15px;
    }
    #footer li:last-child {
        margin-right: 0;
        border-right: 0;
    }
    #footer li:last-child a {
        margin-right: 0;
    }
    #footer a {
        color: #666;
        text-decoration: underline;
        background-image: none;
    }
    #footer a:hover {
        color: #858585;
    }
    #footer .footer-left {
        height: 20px;
        vertical-align: middle;
        line-height: 20px;
    }
    @media (min-width: 39rem) {
        #footer {
            display: flex;
            flex-flow: row wrap;
            justify-content: space-between;
            align-items: center;
            align-content: center;
            margin-bottom: 20px;
        }
        #footer .footer-left {
            align-self: flex-start;
            margin-right: 20px;
        }
        #footer .footer-right {
            align-self: flex-end;
        }
    }
</style>
