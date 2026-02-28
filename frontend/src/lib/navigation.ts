import navigationConfig from "@/config/navigation.json";

export interface NavItem {
  label: string;
  href?: string;
  icon: string;
  children?: NavItem[];
  enabled?: boolean;
  isDisplay?: boolean;
  tags?: string[];
}

interface NavigationConfig {
  navItems: NavItem[];
}

const processNavItem = (item: (typeof navigationConfig.navItems)[0]): NavItem | null => {
  if (item.isDisplay === false) return null;
  const processedItem: NavItem = {
    ...item,
    enabled: item.enabled !== undefined ? item.enabled : true,
  };
  if (item.children) {
    const processedChildren = item.children
      .map((child) => processNavItem(child as (typeof navigationConfig.navItems)[0]))
      .filter((c): c is NavItem => c !== null);
    if (processedChildren.length > 0) processedItem.children = processedChildren;
  }
  return processedItem;
};

export const getNavigationItems = async (): Promise<NavItem[]> => {
  const processedItems: NavItem[] = [];
  for (const item of (navigationConfig as NavigationConfig).navItems) {
    const processed = processNavItem(item);
    if (processed) processedItems.push(processed);
  }
  return processedItems;
};

export const getBadgeStyle = (tag: string): { bgColor: string; textColor: string; label: string } => {
  switch (tag.toLowerCase()) {
    case "wip":
      return { bgColor: "bg-yellow-500", textColor: "text-yellow-100", label: "WIP" };
    case "new":
      return { bgColor: "bg-green-500", textColor: "text-green-100", label: "NEW" };
    case "beta":
      return { bgColor: "bg-blue-500", textColor: "text-blue-100", label: "BETA" };
    default:
      return { bgColor: "bg-gray-500", textColor: "text-gray-100", label: tag.toUpperCase() };
  }
};

export { navigationConfig };
