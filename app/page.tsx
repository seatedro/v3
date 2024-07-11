import { BlogPosts } from 'app/components/posts'
import WorkProjectsSection from './components/work'

export default function Page() {
  return (
    <section>
      <h1 className="mb-4 text-2xl font-semibold tracking-tighter">
        rohit prakash
      </h1>
      <div className="mb-4 flex flex-row gap-2 items-center">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-4">
          <path strokeLinecap="round" strokeLinejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
          <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z" />
        </svg>
        <h3 className="text-center align-middle">
          nyc, usa
        </h3>
      </div>
      <div className="mb-4 flex flex-row gap-2 items-center">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-4">
          <path strokeLinecap="round" strokeLinejoin="round" d="M2.25 21h19.5m-18-18v18m10.5-18v18m6-13.5V21M6.75 6.75h.75m-.75 3h.75m-.75 3h.75m3-6h.75m-.75 3h.75m-.75 3h.75M6.75 21v-3.375c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21M3 3h12m-.75 4.5H21m-3.75 3.75h.008v.008h-.008v-.008Zm0 3h.008v.008h-.008v-.008Zm0 3h.008v.008h-.008v-.008Z" />
        </svg>
        <h3 className="text-center align-middle">
          sre intern @ <a rel="noopener noreferrer" target="_blank" href="https://nbcuniversal.com">nbcuniversal</a>
        </h3>
      </div>
      <p className="mb-4">
        {`currently a student (mscs @ nyu). i love learning new things and 
        building with it. i love programming languages and neovim. currently
        trying to beat the final boss in elden ring but failing miserably.`}
      </p>
      <div className="my-8">
        <WorkProjectsSection />
      </div>
      <div className="my-8">
        <BlogPosts />
      </div>
    </section >
  )
}
