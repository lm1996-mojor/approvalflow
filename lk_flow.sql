/*
 Navicat Premium Data Transfer

 Source Server         : 本机
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : lk_flow

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 29/05/2023 15:11:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comparison_classify
-- ----------------------------
DROP TABLE IF EXISTS `comparison_classify`;
CREATE TABLE `comparison_classify`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `classify_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类名',
  `classify_char` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类标识（int、float、str）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '比较符分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comparison_operators
-- ----------------------------
DROP TABLE IF EXISTS `comparison_operators`;
CREATE TABLE `comparison_operators`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `classify_id` bigint NULL DEFAULT NULL COMMENT '比较符分类id',
  `comparison_operator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '比较符',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '比较符信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comparison_value
-- ----------------------------
DROP TABLE IF EXISTS `comparison_value`;
CREATE TABLE `comparison_value`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `cond_detail_info_id` bigint NULL DEFAULT NULL COMMENT '条件详细信息id',
  `comp_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '比较值',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '比较值信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for condition_detail_groups
-- ----------------------------
DROP TABLE IF EXISTS `condition_detail_groups`;
CREATE TABLE `condition_detail_groups`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `point_id` bigint NULL DEFAULT NULL COMMENT '节点id',
  `group_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分组名称',
  `group_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分组编号',
  `pass_ask` tinyint NULL DEFAULT 2 COMMENT '通过要求（1 满足全部 2 其中之一）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '条件详细信息分组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for condition_detail_info
-- ----------------------------
DROP TABLE IF EXISTS `condition_detail_info`;
CREATE TABLE `condition_detail_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `groups_id` bigint NULL DEFAULT NULL COMMENT '分组id',
  `ctl_id` bigint NULL DEFAULT NULL COMMENT '控件id',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '条件详细信息表(条件和控件关系表)' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for control_info
-- ----------------------------
DROP TABLE IF EXISTS `control_info`;
CREATE TABLE `control_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父级id（关联流程主体id和规则表id）',
  `tab_id` bigint NULL DEFAULT NULL COMMENT '标签id',
  `owner_type` tinyint NULL DEFAULT 1 COMMENT '控件所属类型（1 流程主体 2 流程规则）',
  `cn_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '控件中文名（唯一）',
  `en_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '控件英文名（唯一）',
  `ctl_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '控件编码（唯一）',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否开启（1 开启 2 禁用）',
  `required` tinyint(1) NULL DEFAULT NULL COMMENT '控件值是否必填（1 是 2 否）',
  `field_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '控件数据库表列名（唯一）',
  `component_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '控件类型',
  `value_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '控件值类型',
  `props` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '控件属性',
  `order_no` int NULL DEFAULT NULL COMMENT '控件排序',
  `is_default` tinyint(1) NULL DEFAULT NULL COMMENT '是否为默认控件（1 是 2 否）',
  `is_custom` tinyint(1) NULL DEFAULT NULL COMMENT '是否自建控件（1 是 2 否）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '控件表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ctl_tab
-- ----------------------------
DROP TABLE IF EXISTS `ctl_tab`;
CREATE TABLE `ctl_tab`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `tab_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签名称',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '控件标签（条件要素标签）' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ctl_value_info
-- ----------------------------
DROP TABLE IF EXISTS `ctl_value_info`;
CREATE TABLE `ctl_value_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `approval_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '审批编号',
  `ctl_id` bigint NULL DEFAULT NULL COMMENT '控件id',
  `ctl_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '控件值信息',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '控件值信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for examine_rule
-- ----------------------------
DROP TABLE IF EXISTS `examine_rule`;
CREATE TABLE `examine_rule`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_rule_id` bigint NULL DEFAULT NULL COMMENT '父级规则id',
  `rule_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '规则编号',
  `rule_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '规则名称',
  `rule_explain` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '规则说明',
  `rule_type` int NULL DEFAULT NULL COMMENT '规则类型（参与者形式规则、流程规则、节点规则、异常规则）',
  `rule_level` int NULL DEFAULT NULL COMMENT '规则等级（1为最大，依次递增，权限依次递减）',
  `rule_action_func_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '规则执行方法名',
  `rule_action_func_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '规则执行方法保存路径',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程规则主体表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for panel_point
-- ----------------------------
DROP TABLE IF EXISTS `panel_point`;
CREATE TABLE `panel_point`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `process_id` bigint NULL DEFAULT NULL COMMENT '流程id',
  `point_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点名称',
  `point_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点编号',
  `scenario` tinyint NULL DEFAULT NULL COMMENT '节点使用场景（待定字段）',
  `point_type` tinyint NULL DEFAULT NULL COMMENT '节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）',
  `examine_type` tinyint NULL DEFAULT NULL COMMENT '审批形式（1 依次审批 2 会签 3 或签）',
  `priority` int NULL DEFAULT NULL COMMENT '条件节点优先级(1为最大，优先级依次递减,默认条件优先级最低)',
  `condition_type` tinyint NULL DEFAULT 1 COMMENT '条件创建类型（1 自建 2 默认条件）',
  `only_web_use_previous_step` bigint NULL DEFAULT NULL COMMENT '上一节点id（仅限前端遍历使用）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 212 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for participant
-- ----------------------------
DROP TABLE IF EXISTS `participant`;
CREATE TABLE `participant`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `obj_id` bigint NULL DEFAULT NULL COMMENT '参与者id',
  `point_value_id` bigint NULL DEFAULT NULL COMMENT '节点值id',
  `order_no` int NULL DEFAULT NULL COMMENT '参与者顺序',
  `approval_result` tinyint NULL DEFAULT NULL COMMENT '审批结果（1 同意 2退回 3驳回 4审批中 5待执行 6无操作）',
  `opinions` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '审批意见',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '参与者信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for participant_format
-- ----------------------------
DROP TABLE IF EXISTS `participant_format`;
CREATE TABLE `participant_format`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `format_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '形式名称',
  `format_explain` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '形式说明',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '参与者形式信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for point_participant_rel
-- ----------------------------
DROP TABLE IF EXISTS `point_participant_rel`;
CREATE TABLE `point_participant_rel`  (
  `id` bigint NOT NULL COMMENT '主键id',
  `point_id` bigint NULL DEFAULT NULL COMMENT '节点id',
  `participant_format_id` bigint NULL DEFAULT NULL COMMENT '参与者形式id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程节点与参与者形式关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for point_rel_desc
-- ----------------------------
DROP TABLE IF EXISTS `point_rel_desc`;
CREATE TABLE `point_rel_desc`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `point_id` bigint NULL DEFAULT NULL COMMENT '节点id',
  `previous_step_type` tinyint NULL DEFAULT 0 COMMENT '上一节点类型(1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点)',
  `previous_step` bigint NULL DEFAULT 0 COMMENT '上一节点',
  `next_step` bigint NULL DEFAULT 0 COMMENT '下一节点',
  `next_step_type` tinyint NULL DEFAULT 0 COMMENT '下一节点类型1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 219 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程节点关系描述表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for point_value
-- ----------------------------
DROP TABLE IF EXISTS `point_value`;
CREATE TABLE `point_value`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `approval_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '审批编号（32位）',
  `point_id` bigint NULL DEFAULT NULL COMMENT '节点id',
  `point_rate` tinyint NULL DEFAULT NULL COMMENT '节点进度（1 同意 2退回 3驳回 4审批中 5待执行）',
  `next_step` bigint NULL DEFAULT 0 COMMENT '下一节点',
  `next_step_type` tinyint NULL DEFAULT 0 COMMENT '下一节点类型1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '节点结果信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for process
-- ----------------------------
DROP TABLE IF EXISTS `process`;
CREATE TABLE `process`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `group_id` bigint NULL DEFAULT NULL COMMENT '流程分组id',
  `flow_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '流程名',
  `flow_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '流程编号',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '流程图标',
  `illustrate` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '流程说明',
  `subsidiary_form_id` bigint NULL DEFAULT NULL COMMENT '流程附属表单（表单来源为外部来源时才有值）',
  `form_source` tinyint NULL DEFAULT NULL COMMENT '流程表单来源（流程自建、外部来源）',
  `process_status` tinyint NULL DEFAULT NULL COMMENT '流程状态(1 已发布 2 未发布)',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程主体信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for process_groups
-- ----------------------------
DROP TABLE IF EXISTS `process_groups`;
CREATE TABLE `process_groups`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `client_id` bigint NULL DEFAULT 0 COMMENT '租户id',
  `app_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用编码',
  `business_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务编码',
  `group_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分组名称',
  `order_no` int NULL DEFAULT NULL COMMENT '分组排序',
  `is_default` tinyint NULL DEFAULT 2 COMMENT '是否默认（1 是 2 否）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程分组表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for process_obj_rule_rel
-- ----------------------------
DROP TABLE IF EXISTS `process_obj_rule_rel`;
CREATE TABLE `process_obj_rule_rel`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `rule_id` bigint NULL DEFAULT NULL COMMENT '规则id',
  `process_obj_id` bigint NULL DEFAULT NULL COMMENT '路程对象id',
  `owner_type` int NULL DEFAULT NULL COMMENT '关系所属对象类型（1 参与者、2 流程主体、3 节点）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程对象和规则关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for process_value
-- ----------------------------
DROP TABLE IF EXISTS `process_value`;
CREATE TABLE `process_value`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `app_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用编码',
  `business_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务编码',
  `client_id` int NULL DEFAULT 0 COMMENT '租户id',
  `process_id` bigint NULL DEFAULT NULL COMMENT '流程id',
  `approval_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '审批编号(32位)',
  `process_rate` tinyint NULL DEFAULT NULL COMMENT '流程结果进度（1 同意 2退回 3驳回 4审批中 5待执行）',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流程结果信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for subset
-- ----------------------------
DROP TABLE IF EXISTS `subset`;
CREATE TABLE `subset`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `process_id` bigint NULL DEFAULT NULL COMMENT '流程id',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父级子集id',
  `cn_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '子集中文名（唯一）',
  `en_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '子集英文名（唯一）',
  `sub_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '子集编码（唯一）',
  `field_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '子集数据库表名（唯一）',
  `enable` tinyint NULL DEFAULT 1 COMMENT '子集是否开启（1 开启 2 禁用）',
  `is_default` tinyint(1) NULL DEFAULT NULL COMMENT '是否为默认子集（1 是 2 否）',
  `order_no` int NULL DEFAULT NULL COMMENT '子集排序',
  `props` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '子集属性',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '软删标识（有值时代表删除）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '子集信息表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
