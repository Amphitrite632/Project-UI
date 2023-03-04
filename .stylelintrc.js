module.exports = {
    plugins: ["stylelint-prettier"],
    extends: ["stylelint-config-standard", "stylelint-config-recess-order"],
    rules: {
        "color-hex-length": "long",
        "font-family-name-quotes": "always-unless-keyword",
        "font-family-no-missing-generic-family-keyword": [
            true,
            {
                ignoreFontFamilies: "Material Icons",
            },
        ],
        "selector-class-pattern": "^([A-Z]{2,}|[a-z])+([A-Z][a-z]+)*$",
        "selector-id-pattern": "^([A-Z]{2,}|[a-z])+([A-Z][a-z]+)*$",
        "keyframes-name-pattern": "^([A-Z]{2,}|[a-z])+([A-Z][a-z]+)*$",
    },
}
