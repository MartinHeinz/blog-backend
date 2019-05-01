<template>
    <li><a :href="url"> {{ value }}</a></li>
</template>

<script>
export default {
    name: 'BaseMenuItem',
    props: {
        url: String,
        value: String,
    },
};
</script>

<style>

</style>
