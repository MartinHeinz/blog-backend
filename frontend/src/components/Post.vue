<template>
    <article class="post" itemscope="" itemtype="http://schema.org/BlogPosting">

        <div class="content" itemprop="articleBody">
            <p>{{ text }}</p>
            <span>
                When \(a \ne 0\), there are two solutions to \(ax^2 + bx + c = 0\) and they are
                $$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$
            </span>
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

<style>

    span {
    }
</style>
