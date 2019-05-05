<template>
    <header>

        <h1 class="posttitle" itemprop="name headline">{{title}}</h1>


        <div class="meta">

            <span class="author" itemprop="author" itemscope="" itemtype="http://schema.org/Person">
                <span itemprop="name">{{author}}</span>
            </span>

            <div class="postdate">
                <time datetime="..." itemprop="datePublished">{{published}}</time>
            </div>

            <div class="article-tag">
                <v-icon>fas fa-tag</v-icon>
                <a class="tag-link" href="/tags/.../">...</a>, <a class="tag-link" href="/tags/.../">...</a>
            </div>

        </div>
    </header>
</template>

<script>
// TODO use Moment JS for datetime="..."
// TODO add parameters for tags
// TODO add css
export default {
    name: 'PostHeader',
    components: { },
    props: {
        title: String,
        author: String,
        published: String,
    },
};
</script>

<style>

</style>
