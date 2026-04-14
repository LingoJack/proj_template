export default function Home() {
  return (
    <div className="max-w-3xl mx-auto">
      <h2 className="text-2xl font-bold text-gray-800 mb-4">Welcome</h2>
      <p className="text-gray-600 mb-6">
        全栈项目模板，基于 Go (Echo + GORM) + React (Vite + Tailwind CSS)。
      </p>
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div className="bg-white rounded-lg border border-gray-200 p-4">
          <h3 className="font-medium text-gray-800 mb-1">Backend</h3>
          <p className="text-sm text-gray-500">Echo v4 + GORM + Wire + Zerolog</p>
        </div>
        <div className="bg-white rounded-lg border border-gray-200 p-4">
          <h3 className="font-medium text-gray-800 mb-1">Frontend</h3>
          <p className="text-sm text-gray-500">React 19 + Vite + Tailwind CSS + Zustand</p>
        </div>
        <div className="bg-white rounded-lg border border-gray-200 p-4">
          <h3 className="font-medium text-gray-800 mb-1">API Docs</h3>
          <p className="text-sm text-gray-500">Swagger / OpenAPI</p>
        </div>
        <div className="bg-white rounded-lg border border-gray-200 p-4">
          <h3 className="font-medium text-gray-800 mb-1">DevOps</h3>
          <p className="text-sm text-gray-500">Docker + docker-compose + CI</p>
        </div>
      </div>
    </div>
  )
}
