
// 用户名校验规则：长度在 3 到 12 之间，必须为字母或数字
import {
    PasswordCharacterKindRuleKey,
    PasswordLengthRuleKey,
    PasswordRuleRequiredKey,
    UsernameCharacterKindRuleKey,
    UsernameLengthRuleKey,
    UsernameRuleRequiredKey
} from '@/locales/locale'

export const validateUsername = () => [
    {
        required: true,
        message: UsernameRuleRequiredKey,
    },
    {
        min: 3,
        max: 12,
        message: UsernameLengthRuleKey,
    },
    {
        pattern: /^[a-zA-Z0-9]+$/,
        message: UsernameCharacterKindRuleKey,
    },
];

// 密码校验规则：长度在 8 到 20 之间，至少包含大写字母、小写字母、数字和特殊字符中的三种
export const validatePassword = () => [
    {
        required: true,
        message: PasswordRuleRequiredKey,
    },
    {
        min: 8,
        max: 20,
        message: PasswordLengthRuleKey,
    },
    {
        validator: (_: any, value: string) => {
            if (!value) {
                return Promise.resolve();
            }
            // 使用正则表达式分别匹配大写字母、小写字母、数字和特殊字符
            const hasUpperCase = /[A-Z]/.test(value);
            const hasLowerCase = /[a-z]/.test(value);
            const hasNumber = /\d/.test(value);
            const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(value);

            const count = [hasUpperCase, hasLowerCase, hasNumber, hasSpecialChar].filter(Boolean).length;

            if (count >= 3) {
                return Promise.resolve();
            }
            return Promise.reject(new Error('密码至少包含大写字母、小写字母、数字和特殊字符中的三种'));
        },
        message: PasswordCharacterKindRuleKey,
    },
];
