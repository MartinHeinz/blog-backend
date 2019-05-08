<template>
    <article class="post" itemscope="" itemtype="http://schema.org/BlogPosting">

        <div class="content" itemprop="articleBody">
            <p>{{ text }}</p>
        </div>
    </article>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Post',
    props: {
        text: String,
    },
    mounted() {
        axios
            .get('http://localhost:1234/api/v1/')
            .then((response) => { (this.info = response); });
    },
};
</script>
