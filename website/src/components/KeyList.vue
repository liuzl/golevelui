<template>
  <div class="key-list-container">
    <el-row :gutter="20">
      <!-- Key List Column -->
      <el-col :xs="24" :sm="10" :md="10" :lg="10" :xl="10">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>Keys</span>
            <el-button
              @click="handleNew"
              style="float: right; padding: 3px 0"
              type="text"
              icon="el-icon-plus"
            >New Key</el-button>
          </div>
          <el-input
            placeholder="Filter by prefix"
            v-model="prefix"
            class="filter-input"
            clearable
          >
            <i slot="prefix" class="el-input__icon el-icon-search"></i>
          </el-input>
          <el-table
            highlight-current-row
            :data="data"
            class="list-table"
            :border="false"
            ref="table"
            :height="tableHeight"
            @row-click="(row) => handleItemClick(row, false)"
          >
            <el-table-column prop="keyName" label="Key"></el-table-column>
            <el-table-column width="60" fixed="right" label="Actions">
              <template slot-scope="scope">
                <el-button
                  @click.stop="handleDelete(scope.row)"
                  size="mini"
                  type="danger"
                  icon="el-icon-delete"
                  circle
                ></el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="foot">
            <el-button
              :disabled="!searchText"
              size="mini"
              icon="el-icon-arrow-down"
              @click="next"
            >Load More</el-button>
            <el-tag size="mini" :type="countType" @click="loadCounts">Total: {{ count }}</el-tag>
          </div>
        </el-card>
      </el-col>

      <!-- Value Editor Column (Desktop) -->
      <el-col :xs="0" :sm="14" :md="14" :lg="14" :xl="14">
        <el-card class="box-card value-editor">
          <div slot="header" class="clearfix">
            <span>Value</span>
            <el-button
              v-if="currentKey"
              @click="handleUpdate"
              style="float: right; padding: 3px 0"
              type="text"
              icon="el-icon-check"
            >Save</el-button>
          </div>
          <div class="editor-toolbar">
             <el-select v-model="format" placeholder="Format" size="mini" clearable>
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              ></el-option>
            </el-select>
          </div>
          <el-input
            :rows="editorRows"
            type="textarea"
            resize="none"
            placeholder="Select a key to view its value"
            v-model="currentValue"
            class="value-textarea"
          ></el-input>
        </el-card>
      </el-col>
    </el-row>

    <!-- Value Editor Drawer (Mobile) -->
    <el-drawer
      :visible.sync="drawerVisible"
      :with-header="false"
      direction="rtl"
      size="85%"
    >
      <el-card class="box-card value-editor-drawer">
        <div slot="header" class="clearfix">
          <span>Value</span>
           <el-button
            v-if="currentKey"
            @click="handleUpdate"
            style="float: right; padding: 3px 0"
            type="text"
            icon="el-icon-check"
          >Save</el-button>
        </div>
        <div class="editor-toolbar">
           <el-select v-model="format" placeholder="Format" size="mini" clearable>
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </div>
        <el-input
          :rows="editorRows"
          type="textarea"
          resize="none"
          placeholder="Select a key to view its value"
          v-model="currentValue"
          class="value-textarea"
        ></el-input>
      </el-card>
    </el-drawer>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from "vue-property-decorator";
import {
  Table,
  TableColumn,
  Message,
  Input,
  Row,
  Col,
  Button,
  Select,
  Option,
  Tag,
  MessageBox,
  Card,
  Drawer
} from "element-ui";
import { keys, keyInfo, keyDelete, keyUpdate, keysCount } from "@/api/golevelui";

interface Item {
  keyName: string;
}

@Component({
  components: {
    ElTable: Table,
    ElTableColumn: TableColumn,
    ElInput: Input,
    ElRow: Row,
    ElCol: Col,
    ElButton: Button,
    ElSelect: Select,
    ElOption: Option,
    ElTag: Tag,
    ElCard: Card,
    ElDrawer: Drawer
  }
})
export default class List extends Vue {
  @Prop({ default: "" })
  private db!: string;

  private data: Item[] = [];
  private prefix = "";
  private tableHeight = 400;
  private currentValue = "";
  private currentKey = "";
  private format = "";
  private searchText = "";
  private countSearchText = "";
  private countType = "success";
  private count = 0;
  private countIsTrue = false;
  private countLock = false;
  
  private isMobile = false;
  private drawerVisible = false;

  private options = [{
    label: "Json",
    value: "Json"
  }];

  get editorRows() {
    // Adjust editor height based on viewport
    return Math.floor((window.innerHeight - 280) / 21);
  }

  // --- Lifecycle & Resize ---
  mounted() {
    this.handleResize();
    window.addEventListener("resize", this.handleResize);
    this.$nextTick(() => {
      this.tableHeight = (this.$el as HTMLElement).clientHeight - 160;
    });
  }

  beforeDestroy() {
    window.removeEventListener("resize", this.handleResize);
  }

  handleResize() {
    this.isMobile = window.innerWidth < 768; // Element UI's sm breakpoint
    this.tableHeight = (this.$el as HTMLElement).clientHeight - 160;
  }

  // --- Data Loading ---
  created() {
    this.loadKeys();
    this.loadCounts();
  }

  loadKeys() {
    if (!this.db) return;
    keys({
      db: this.db,
      prefix: this.prefix,
      searchText: this.searchText
    }).then(res => {
      this.searchText = res.data.IsPart ? res.data.SearchText : "";
      const newItems = res.data.Items.map((item: string) => ({ keyName: item }));
      this.data = this.data.concat(newItems);
    });
  }

  loadCounts() {
    if (this.countLock || this.countIsTrue || !this.db) return;
    this.countLock = true;
    keysCount({
      db: this.db,
      prefix: this.prefix,
      searchText: this.countSearchText
    })
      .then(res => {
        this.countSearchText = res.data.LastKey;
        this.count += res.data.Count;
        this.countType = res.data.IsTrue ? "success" : "warning";
        this.countIsTrue = res.data.IsTrue;
      })
      .finally(() => {
        this.countLock = false;
      });
  }

  next() {
    this.loadKeys();
  }

  @Watch("prefix")
  onPrefixChange() {
    this.data = [];
    this.searchText = "";
    this.count = 0;
    this.countSearchText = "";
    this.countIsTrue = false;
    this.loadKeys();
    this.loadCounts();
  }

  // --- Item Interaction ---
  handleItemClick(row: Item, isNew = false) {
    if (!this.db) return Message.error("Invalid DB name");
    this.currentKey = row.keyName;

    if (this.isMobile) {
      this.drawerVisible = true;
    }

    if (isNew) {
      this.currentValue = "";
      Message.info("Enter a value and click 'Save' to create the new key.");
      return;
    }

    keyInfo({ db: this.db, key: row.keyName }).then(res => {
      let formatValue = res.data.value;
      if (this.format === "Json") {
        formatValue = this.formatValueToJson(res.data.value);
      }
      this.currentValue = formatValue;
    });
  }

  handleNew() {
    MessageBox.prompt("Please enter the new key", "New Key", {
      confirmButtonText: "OK",
      cancelButtonText: "Cancel"
    })
      .then(({ value }) => {
        if (value) {
          const newKey = { keyName: value };
          this.data.unshift(newKey);
          this.count++;
          this.handleItemClick(newKey, true);
        }
      })
      .catch(() => {
        Message.info("Canceled");
      });
  }

  handleUpdate() {
    if (!this.currentKey) return Message.info("Please select a key");
    
    let valueToSave = this.currentValue;
    if (this.format === 'Json') {
        valueToSave = this.formatJsonToValue(this.currentValue);
    }

    keyUpdate({ db: this.db, key: this.currentKey, value: valueToSave })
      .then(res => {
        if (res.data.Success) {
          Message.success("Update successful!");
          if (this.isMobile) {
            this.drawerVisible = false;
          }
        } else {
          Message.error("Update failed!");
        }
      })
      .catch(err => {
        Message.error(`Update failed: ${err}`);
      });
  }

  handleDelete(row: Item) {
    MessageBox.confirm("This will permanently delete the key. Continue?", "Warning", {
      confirmButtonText: "OK",
      cancelButtonText: "Cancel",
      type: "warning"
    })
      .then(() => {
        keyDelete({ db: this.db, key: row.keyName }).then(res => {
          if (res.data.Success) {
            Message.success("Delete successful!");
            this.data.splice(this.data.indexOf(row), 1);
            this.count--;
            if (this.currentKey === row.keyName) {
                this.currentKey = "";
                this.currentValue = "";
            }
          } else {
            Message.error("Delete failed!");
          }
        });
      })
      .catch(() => {
        Message.info("Delete canceled");
      });
  }

  // --- Formatting ---
  @Watch("format")
  formatChange(now: string, old: string) {
    if (now === "Json") {
      this.currentValue = this.formatValueToJson(this.currentValue);
    } else if (old === "Json") {
      this.currentValue = this.formatJsonToValue(this.currentValue);
    }
  }

  formatValueToJson(value: string): string {
    try {
      return JSON.stringify(JSON.parse(value), null, 2);
    } catch (e) {
      return value;
    }
  }

  formatJsonToValue(json: string): string {
    try {
      return JSON.stringify(JSON.parse(json));
    } catch (e) {
      return json;
    }
  }
}
</script>

<style scoped>
.key-list-container {
  height: 100%;
}
.el-row, .el-col {
  height: 100%;
}
.box-card {
  height: 100%;
  display: flex;
  flex-direction: column;
}
::v-deep .el-card__header {
  flex-shrink: 0;
}
::v-deep .el-card__body {
  flex-grow: 1;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}
.filter-input {
  margin-bottom: 15px;
}
.list-table {
  width: 100%;
  flex-grow: 1;
}
.foot {
  padding-top: 15px;
  text-align: center;
  flex-shrink: 0;
}
.foot .el-tag {
  margin-left: 10px;
  cursor: pointer;
}
.value-editor .el-card__body {
  padding: 0;
}
.editor-toolbar {
  padding: 10px;
  background-color: #f9fafc;
  border-bottom: 1px solid #ebeef5;
}
.value-textarea {
  flex-grow: 1;
}
::v-deep .value-textarea textarea {
  height: 100%;
  border: none;
  border-radius: 0;
}
.value-editor-drawer {
    height: 100vh;
}
::v-deep .el-table__body tr.current-row > td {
  background-color: #ecf5ff !important;
}
</style>
