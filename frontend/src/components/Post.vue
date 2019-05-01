<template>
    <div class="content index py4">
        <h1>{{ info }}</h1>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Post',
    props: {
        info: String,
    },
    mounted() {
        axios
            .get('http://localhost:1234/api/v1/')
            .then((response) => { (this.info = response); });
    },
};
</script>
