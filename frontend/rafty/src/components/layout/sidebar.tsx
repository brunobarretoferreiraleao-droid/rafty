export function Sidebar() {
  return (
    <aside className="hidden w-64 shrink-0 border-r border-white/10 p-6 lg:block">
      <div className="space-y-3">
        <div className="rounded-2xl bg-white/3 p-4">
          Profile
        </div>

        <div className="rounded-2xl bg-white/3 p-4">
          Saved
        </div>

        <div className="rounded-2xl bg-white/3 p-4">
          Friends
        </div>
      </div>
    </aside>
  );
}