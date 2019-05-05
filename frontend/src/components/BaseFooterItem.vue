<template>
    <li><a :href="url"> {{ value }}</a></li>
</template>

<script>
export default {
    name: 'BaseFooterItem',
    props: {
        url: String,
        value: String,
    },
};
</script>

<style>

</style>
