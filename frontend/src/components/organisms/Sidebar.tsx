"use client";

import { app, env, ui } from "@/config";
import { getUserProfile, logout, UserProfile } from "@/lib/auth";
import { getBadgeStyle, getNavigationItems, NavItem } from "@/lib/navigation";
import {
  AcademicCapIcon,
  ArrowRightEndOnRectangleIcon,
  Bars3Icon,
  BookOpenIcon,
  CalendarDaysIcon,
  ChartBarIcon,
  ChevronDownIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  ClipboardDocumentCheckIcon,
  ClipboardDocumentListIcon,
  ComputerDesktopIcon,
  DocumentDuplicateIcon,
  DocumentTextIcon,
  HomeIcon,
  KeyIcon,
  LinkIcon,
  NewspaperIcon,
  PhotoIcon,
  ServerIcon,
  ShieldCheckIcon,
  UserGroupIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import ConfirmDialog from "./ConfirmDialog";

const iconMap: Record<string, React.ComponentType<{ className?: string }>> = {
  HomeIcon,
  ChartBarIcon,
  ServerIcon,
  UserGroupIcon,
  DocumentTextIcon,
  BookOpenIcon,
  AcademicCapIcon,
  ClipboardDocumentListIcon,
  DocumentDuplicateIcon,
  CalendarDaysIcon,
  ComputerDesktopIcon,
  ClipboardDocumentCheckIcon,
  ShieldCheckIcon,
  KeyIcon,
  PhotoIcon,
  NewspaperIcon,
  LinkIcon,
};

export default function Sidebar() {
  const router = useRouter();
  const [collapsed, setCollapsed] = useState(false);
  const [width, setWidth] = useState(280);
  const [expandedItems, setExpandedItems] = useState<string[]>([]);
  const [userProfile, setUserProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(false);
  const [showLogoutConfirm, setShowLogoutConfirm] = useState(false);
  const [navItems, setNavItems] = useState<NavItem[]>([]);
  const [mobile, setMobile] = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);
  const sidebarRef = useRef<HTMLDivElement>(null);
  const resizing = useRef(false);

  useEffect(() => {
    const check = () => setMobile(window.innerWidth < 768);
    check();
    window.addEventListener("resize", check);
    return () => window.removeEventListener("resize", check);
  }, []);

  useEffect(() => {
    getNavigationItems().then(setNavItems).catch(console.error);
  }, []);

  useEffect(() => {
    getUserProfile().then(setUserProfile).catch(console.error);
  }, []);

  const handleLogout = async () => {
    try {
      setLoading(true);
      await logout();
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  };

  const toggleExpand = (label: string) => {
    setExpandedItems((prev) =>
      prev.includes(label) ? prev.filter((l) => l !== label) : [...prev, label]
    );
  };

  const goTo = (href: string) => {
    if (mobile) setMobileOpen(false);
    router.push(href);
  };

  const renderNavItem = (item: NavItem, level = 0) => {
    const Icon = iconMap[item.icon];
    const hasChildren = item.children && item.children.length > 0;
    const isExpanded = expandedItems.includes(item.label);
    const enabled = item.enabled !== false;
    const textColor =
      level === 0 ? "text-white" : level === 1 ? "text-blue-100" : "text-blue-200";

    if (hasChildren) {
      return (
        <div key={item.label}>
          <button
            type="button"
            onClick={() => toggleExpand(item.label)}
            className={`w-full flex items-center py-2 rounded-xl hover:bg-blue-600/80 font-medium ${textColor} ${collapsed && !mobile ? "justify-center" : ""}`}
            style={{
              paddingLeft: collapsed && !mobile ? "0.75rem" : `${0.75 + level * 1.5}rem`,
              paddingRight: "0.75rem",
            }}
          >
            {Icon && <Icon className="h-5 w-5 mr-3 flex-shrink-0" />}
            {(!collapsed || mobile) && (
              <>
                <span className="truncate flex-1 text-left flex items-center">
                  {item.label}
                  {item.tags?.[0] && (
                    <span
                      className={`ml-2 px-1.5 py-0.5 text-xs rounded ${getBadgeStyle(item.tags[0]).bgColor} ${getBadgeStyle(item.tags[0]).textColor}`}
                    >
                      {getBadgeStyle(item.tags[0]).label}
                    </span>
                  )}
                </span>
                {isExpanded ? (
                  <ChevronDownIcon className="h-4 w-4 ml-2" />
                ) : (
                  <ChevronRightIcon className="h-4 w-4 ml-2" />
                )}
              </>
            )}
          </button>
          <div
            className={`overflow-hidden transition-all duration-300 ${(!collapsed || mobile) && isExpanded ? "max-h-96 opacity-100" : "max-h-0 opacity-0"}`}
          >
            {(!collapsed || mobile) && isExpanded && (
              <div className="mt-2 space-y-1">
                {item.children!.map((c) => renderNavItem(c, level + 1))}
              </div>
            )}
          </div>
        </div>
      );
    }

    return (
      <button
        key={item.label}
        type="button"
        onClick={() => item.href && goTo(item.href)}
        disabled={!enabled}
        className={`w-full flex items-center py-2 rounded-xl font-medium ${collapsed && !mobile ? "justify-center" : ""} ${enabled ? `${textColor} hover:bg-blue-600/80` : "text-white/50 cursor-not-allowed"}`}
        style={{
          paddingLeft: collapsed && !mobile ? "0.75rem" : `${0.75 + level * 1.5}rem`,
          paddingRight: "0.75rem",
        }}
      >
        {Icon && <Icon className="h-5 w-5 mr-3 flex-shrink-0" />}
        {(!collapsed || mobile) && (
          <span className="truncate flex items-center">
            {item.label}
            {item.tags?.[0] && (
              <span
                className={`ml-2 px-1.5 py-0.5 text-xs rounded ${getBadgeStyle(item.tags[0]).bgColor} ${getBadgeStyle(item.tags[0]).textColor}`}
              >
                {getBadgeStyle(item.tags[0]).label}
              </span>
            )}
          </span>
        )}
      </button>
    );
  };

  const sidebarContent = (
    <>
      <div className="flex items-center h-16 px-4 border-b border-white/20 overflow-hidden">
        {collapsed && !mobile ? (
          <span className="text-lg font-bold text-white">WP</span>
        ) : (
          <div className="flex items-center space-x-3">
            <img src={ui.logo.src} alt={ui.logo.alt} className="h-10 w-auto" />
            <span className="text-lg font-bold text-white whitespace-nowrap">{app.name}</span>
          </div>
        )}
      </div>
      {env.isDevelopment && (
        <div className="px-4 py-2 bg-yellow-500/20 text-yellow-800 text-center text-xs font-medium">
          {collapsed ? "Dev" : "Development"}
        </div>
      )}
      <nav className="flex-1 px-2 py-4 space-y-2 overflow-y-auto">
        {navItems.map((item) => renderNavItem(item))}
      </nav>
      <div className="mt-auto px-4 py-4 border-t border-white/20">
        {userProfile ? (
          <div className="space-y-3">
            <div className={`flex items-center space-x-3 ${collapsed && !mobile ? "justify-center" : ""}`}>
              <div className="h-10 w-10 rounded-full bg-black/40 flex items-center justify-center flex-shrink-0">
                <span className="text-sm font-semibold text-white">
                  {userProfile.user_fullname
                    .split(" ")
                    .map((n) => n.charAt(0))
                    .join("")
                    .toUpperCase()
                    .slice(0, 2)}
                </span>
              </div>
              {(!collapsed || mobile) && (
                <div className="flex-1 min-w-0">
                  <p className="text-sm font-medium text-white truncate">{userProfile.user_fullname}</p>
                  <p className="text-xs text-white/70 truncate">@{userProfile.user_name}</p>
                </div>
              )}
            </div>
            <button
              type="button"
              onClick={() => setShowLogoutConfirm(true)}
              disabled={loading}
              className={`flex items-center justify-center rounded-lg text-white/70 hover:text-white hover:bg-white/10 text-xs py-1.5 w-full ${collapsed && !mobile ? "mx-auto" : ""}`}
            >
              {loading ? (
                <span className="animate-spin h-3 w-3 border-2 border-white border-t-transparent rounded-full" />
              ) : (
                <>
                  <ArrowRightEndOnRectangleIcon className="h-4 w-4" />
                  {(!collapsed || mobile) && <span className="ml-1">Keluar</span>}
                </>
              )}
            </button>
          </div>
        ) : (
          <Link
            href="/auth/login"
            className="flex items-center justify-center rounded-xl bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 w-full"
          >
            {collapsed && !mobile ? "L" : "Login"}
          </Link>
        )}
      </div>
      <ConfirmDialog
        isOpen={showLogoutConfirm}
        onClose={() => setShowLogoutConfirm(false)}
        onConfirm={async () => {
          setShowLogoutConfirm(false);
          await handleLogout();
        }}
        title="Konfirmasi Logout"
        message="Apakah Anda yakin ingin keluar? Anda perlu login kembali untuk mengakses aplikasi."
        confirmText="Ya, Logout"
        cancelText="Batal"
        type="warning"
      />
    </>
  );

  if (mobile) {
    return (
      <>
        <button
          type="button"
          onClick={() => setMobileOpen((o) => !o)}
          className="fixed top-4 left-4 z-50 p-2 bg-white rounded-lg shadow border border-gray-200"
          aria-label="Menu"
        >
          {mobileOpen ? <XMarkIcon className="h-6 w-6 text-gray-700" /> : <Bars3Icon className="h-6 w-6 text-gray-700" />}
        </button>
        {mobileOpen && (
          <div
            className="fixed inset-0 z-40 bg-black/50"
            onClick={() => setMobileOpen(false)}
            aria-hidden
          >
            <div
              ref={sidebarRef}
              className="fixed left-0 top-0 h-full w-80 max-w-[85vw] bg-gradient-to-b from-blue-800 to-blue-900 shadow-xl"
              style={{ background: `linear-gradient(to top, ${ui.theme.gradient.from}, ${ui.theme.gradient.to})` }}
              onClick={(e) => e.stopPropagation()}
            >
              {sidebarContent}
            </div>
          </div>
        )}
      </>
    );
  }

  return (
    <aside
      ref={sidebarRef}
      style={{
        width: collapsed ? 64 : width,
        background: `linear-gradient(to top, ${ui.theme.gradient.from}, ${ui.theme.gradient.to})`,
      }}
      className="relative flex flex-col min-h-screen border-r shadow-sm transition-[width] duration-300"
    >
      <button
        type="button"
        className="absolute -right-4 top-4 z-10 w-8 h-8 bg-white border border-gray-300 rounded-full flex items-center justify-center shadow hover:bg-gray-50"
        onClick={() => setCollapsed((c) => !c)}
        aria-label={collapsed ? "Expand" : "Collapse"}
        style={{ left: collapsed ? 48 : width - 16 }}
      >
        {collapsed ? <ChevronRightIcon className="h-4 w-4 text-gray-700" /> : <ChevronLeftIcon className="h-4 w-4 text-gray-700" />}
      </button>
      {sidebarContent}
    </aside>
  );
}
