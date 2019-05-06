<template>
    <article class="post" itemscope="" itemtype="http://schema.org/BlogPosting">
    <header>

        <h1 class="posttitle" itemprop="name headline">{{title}}</h1>


        <div class="meta">

            <span class="author" itemprop="author" itemscope="" itemtype="http://schema.org/Person">
                <span itemprop="name">{{author}}</span>
            </span>

            <div class="postdate">
                <time :datetime="published | formatDate" itemprop="datePublished">{{published}}</time>
            </div>

            <div class="article-tag">
                <v-icon>fas fa-tag</v-icon>
                <a class="tag-link" href="/tags/.../">...</a>, <a class="tag-link" href="/tags/.../">...</a>
            </div>

        </div>
    </header>
    </article>
</template>

<script>
// TODO add parameters for tags
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
    a:hover {
        box-shadow: inset 0 0px 0 white, inset 0 -2px 0 #d480aa;
    }

    a {
        color: #c9cacc;
        text-decoration: none;
        background-image: linear-gradient(transparent, transparent 5px, #c9cacc 5px, #c9cacc);
        background-position: bottom;
        background-size: 100% 6px;
        background-repeat: repeat-x;
    }

    article header .posttitle {
        margin-top: 0;
        margin-bottom: 0;
        text-transform: none;
        font-size: 1.5em;
        line-height: 1.25;
    }

    .posttitle {
        text-transform: none;
        font-size: 1.5em;
        line-height: 1.25;
    }

    article header .meta {
        margin-top: 0;
        margin-bottom: 1rem;
    }
    article header .meta * {
        color: #ccc;
        font-size: 0.85rem;
    }

    article header .author {
        text-transform: uppercase;
        letter-spacing: 0.01em;
        font-weight: 700;
    }

    article header .postdate {
        display: inline;
    }

    .article-tag .tag-link {
        background-image: linear-gradient(transparent, transparent 10px, #d480aa 10px, #d480aa);
        background-position: bottom;
        background-size: 100% 6px;
        background-repeat: repeat-x;
    }
    .article-tag .tag-link:before {
        content: "#";
    }

    @media (min-width: 480px) {
        .article-tag {
            display: inline;
        }
        .article-tag:before {
            content: "|";
        }
    }
</style>
