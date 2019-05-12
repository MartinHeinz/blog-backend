<template>
    <div id="share" style="visibility: visible;"
         v-show="share_active">
        <social-sharing :url=path inline-template>
            <ul class="networks">
                <li>
                    <network network="facebook">
                        <i class="fab fa-facebook" ></i>
                    </network>
                </li>
                <li>
                    <network network="linkedin">
                        <i class="fab fa-linkedin"></i>
                    </network>
                </li>
                <li>
                    <network network="pinterest">
                        <i class="fab fa-pinterest"></i>
                    </network>
                </li>
                <li>
                    <network network="reddit">
                        <i class="fab fa-reddit"></i>
                    </network>
                </li>
                <li>
                    <network network="twitter">
                        <i class="fab fa-twitter"></i>
                    </network>
                </li>
                <li>
                    <network network="vk">
                        <i class="fab fa-vk"></i>
                    </network>
                </li>
                <li>
                    <network network="weibo">
                        <i class="fab fa-weibo"></i>
                    </network>
                </li>
                <li>
                    <network network="whatsapp">
                        <i class="fab fa-whatsapp"></i>
                    </network>
                </li>
            </ul>
        </social-sharing>

    </div>
</template>

<script>

export default {
    name: 'SocialSharingList',
    props: {
        items: Array,
        actions: Object,
    },
    mounted() {
        this.$root.$on('toggle_social_sharing', () => {
            this.share_active = !this.share_active;
        });
    },
    data() {
        return {
            share_active: false,
            path: window.location.href,
            fb_active: false,
        };
    },
};
</script>

<style>

    #header-post #share {
        clear: both;
        padding-top: 1rem;
        padding-right: 2rem;
        text-align: right;
    }
    #header-post #share li {
        display: block;
        margin: 0;
    }

    .networks li:hover {
        color: #d480aa;
        cursor: pointer;
    }
</style>
