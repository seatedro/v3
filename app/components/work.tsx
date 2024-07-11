type WorkEx = {
  name: string;
  href: string;
  role: string;
  description: string;
  duration?: string;
};

const WorkProjectsSection = () => {
  const work = [
    {
      name: "NBCUniversal",
      href: "https://nbcuniversal.com",
      role: "swe intern",
      duration: "(aug 2021 - jun 2023)",
      description:
        "shipped go services to prod that collect metrics for aws services.",
    },
    {
      name: "Hiringtek",
      href: "https://hiringtek.com",
      role: "full-stack engineer ",
      duration: "(aug 2021 - jun 2023)",
      description:
        "built an e2e coding platform for online interviews. overhauled performance. improved dx with sota dev tooling.",
    },
  ] satisfies WorkEx[];

  const projects = [
    {
      name: "guam",
      href: "https://github.com/seatedro/guam",
      role: "",
      description:
        "rewrite of lucia-auth in go.",
    },
    {
      name: "mira(me)",
      href: "https://github.com/seatedro/mirame",
      role: "",
      description:
        "my version of the monkeylang interpreter in rust & go.",
    },
    {
      name: "coding.rs",
      href: "https://github.com/seatedro/coding.rs",
      role: "",
      description:
        "juiced up online judge written in rust (btw)",
    },
    {
      name: "all projects →",
      href: "https://github.com/seatedro",
      role: "",
      description: "",
    },
  ] satisfies WorkEx[];

  return (
    <div className="my-8 grid grid-cols-1 gap-8 md:grid-cols-2">
      <section className="text-left">
        <h3 className="mb-6 text-xl font-medium">work</h3>
        {work.map((item, index) => (
          <div key={index}>
            <a
              href={item.href}
              target="_blank"
              className="font-medium underline decoration-gray-400 underline-offset-4 dark:decoration-gray-600"
            >
              {item.name}
            </a>
            <p className="mt-2">{item.role}</p>
            {item.duration && <p className="text-sm text-gray-600">{item.duration}</p>}
            <p className="mt-2 text-neutral-700 dark:text-neutral-300">
              {item.description}
            </p>
            {index !== work.length - 1 && <div className="mt-6"></div>}
          </div>
        ))}
      </section>
      <section className="text-left">
        <h3 className="mb-6 text-xl font-medium">projects</h3>
        {projects.map((item, index) => (
          <div key={index}>
            <a
              href={item.href}
              target="_blank"
              className="font-medium underline decoration-gray-400 underline-offset-4 dark:decoration-gray-600"
            >
              {item.name}
            </a>
            <p className="mt-2">{item.role}</p>
            <p className="mt-2 text-neutral-700 dark:text-neutral-300">
              {item.description}
            </p>
            <div className="mt-6"></div>
          </div>
        ))}
      </section>
    </div>
  );
};

export default WorkProjectsSection;
