<template>
    <a class="tag-link" :href="'/tags/' + url + '/'">{{name}}</a>
</template>

<script>
export default {
    name: 'BaseTag',
    props: {
        url: String,
        name: String,
    },
};
</script>

<style>

    .tag-link {
        background-image: linear-gradient(transparent, transparent 10px, #d480aa 10px, #d480aa);
        background-position: bottom;
        background-size: 100% 6px;
        background-repeat: repeat-x;
    }
    .tag-link:before {
        content: " #";
    }
    .tag-link:after {
        content: ", ";
    }
    .tag-link:last-child:after {
        content: "";
    }

</style>
