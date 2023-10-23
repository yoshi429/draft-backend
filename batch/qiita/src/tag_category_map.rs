use std::collections::HashMap;

pub fn create_map() -> HashMap<&'static str, &'static str> {
    let mut tag_category: HashMap<&str, &str> = HashMap::new();

    tag_category.insert("インフラ", "Infrastructure");
    tag_category.insert("アジャイル", "agile");
    tag_category.insert("ビジネス", "Bussiness");
    tag_category.insert("マーケティング", "Marketing");
    tag_category.insert("kubernetes", "kubernetes");
    tag_category.insert("Docker", "Infrastructure");
    tag_category.insert("要件定義", "System Design");
    tag_category.insert("ワイヤーフレーム", "System Design");
    tag_category.insert("googlecloud", "Infrastructure");
    tag_category.insert("Nuxt", "Frontend");

    tag_category
}
