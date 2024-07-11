import { BlogPosts } from 'app/components/posts'

export const metadata = {
  title: 'Blog',
  description: 'Read my blog.',
}

export default function Page() {
  return (
    <section>
      <h1 className="font-semibold text-2xl mb-8 tracking-tighter">my bl.og</h1>
      <h2 className=''>no posts yet! come back later</h2>
      <BlogPosts />
    </section>
  )
}
